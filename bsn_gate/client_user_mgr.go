package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	"fmt"
	"net"
	"sync"
)

type SClientUserMgr struct {
	M_SApp *SApp

	M_TCPListener      net.Listener
	M_TCPstrListenAddr string
	M_TCPSNotifyClose  *bsn_common.SNotifyClose
	M_TCPSState        *bsn_common.SState

	M_TId2User  map[uint16]*SClientUser
	M_TClientId uint16
	M_MutexUser sync.Mutex
}

func NewSClientUserMgr(vSApp *SApp) (this *SClientUserMgr, err error) {
	GSLog.Debugln("NewSClientUserMgr")
	this = &SClientUserMgr{
		M_SApp:             vSApp,
		M_TCPSNotifyClose:  bsn_common.NewSNotifyClose(),
		M_TCPSState:        bsn_common.NewSState(),
		M_TCPstrListenAddr: fmt.Sprintf(":%v", vSApp.ListenPort()),
		M_TId2User:         make(map[uint16]*SClientUser, 10),
	}

	return
}

func (this *SClientUserMgr) start() (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	err = this.startTCPListen()
	if err != nil {
		return
	}

	return
}

func (this *SClientUserMgr) stop() (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	err = this.stopTCPListen()
	if err != nil {
		return
	}

	GSLog.Mustln("close all user")
	for _, vClientUser := range this.M_TId2User {
		err = vClientUser.close()
		if err != nil {
			return
		}
	}
	this.M_TId2User = nil

	return
}

func (this *SClientUserMgr) stopTCPListen() (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_TCPSState.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		GSLog.Debugln("not listen")
		return
	}

	defer func() {
		if err == nil {
			return
		}
		this.M_TCPSState.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	}()

	err = this.M_TCPListener.Close()
	if err != nil {
		return
	}

	this.M_TCPSNotifyClose.WaitClose()
	return
}

func (this *SClientUserMgr) startTCPListen() (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_TCPSState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		GSLog.Errorln("had listen")
		return
	}

	defer func() {
		if err == nil {
			return
		}
		this.M_TCPSState.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	GSLog.Mustln("listenTCP ", this.M_TCPstrListenAddr)
	this.M_TCPListener, err = net.Listen("tcp", this.M_TCPstrListenAddr)
	if err != nil {
		return
	}

	this.M_TCPSNotifyClose.Clear()
	go this.workerTCPListen()
	return
}

func (this *SClientUserMgr) workerTCPListen() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.M_TCPSNotifyClose.Close()
		this.M_TCPListener = nil
		this.M_TCPSState.Set(bsn_common.CState_Idle)
	}()

	this.M_TCPSState.Change(bsn_common.CState_Op, bsn_common.CState_Runing)

	vbQuit := false
	for !vbQuit {
		vConn, err := this.M_TCPListener.Accept()
		if err != nil {
			GSLog.Errorln(err)
			vbQuit = true
			continue
		}
		this.onTCPClient(vConn)
	}
}

func (this *SClientUserMgr) onTCPClient(vConn net.Conn) (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
			vConn.Close()
		}
	}()

	vClientId := this.genClientId()
	if vClientId == 0 {
		err = errors.New("genClientId fail")
		return
	}

	vstrAddr := vConn.RemoteAddr().String()
	GSLog.Debugf("client connect ClientId=%v addr=%v\n", vClientId, vstrAddr)

	vSClientUser, err := NewSClientUser(this, vConn, vClientId)
	if err != nil {
		return
	}

	err = this.addClient(vSClientUser)
	if err != nil {
		return
	}
	vSClientUser.run()

	return
}

// generate clientid
// if not generate return 0
func (this *SClientUserMgr) genClientId() uint16 {
	for i := 0; i < 100; i++ {
		this.M_TClientId++
		if this.M_TClientId == 0 {
			continue
		}

		vSClientUser := this.getClient(this.M_TClientId)
		if vSClientUser == nil {
			return this.M_TClientId
		}
	}
	return 0
}

func (this *SClientUserMgr) getClient(vClientId uint16) *SClientUser {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	return this.M_TId2User[vClientId]
}

func (this *SClientUserMgr) addClient(vSClientUser *SClientUser) error {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	this.M_TId2User[vSClientUser.ClientId()] = vSClientUser
	return nil
}

func (this *SClientUserMgr) delClient(vTClientId uint16) error {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	delete(this.M_TId2User, vTClientId)
	return nil
}
