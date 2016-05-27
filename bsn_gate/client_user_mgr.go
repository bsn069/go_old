package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	"fmt"
	"net"
	"sync"
	// "sync/atomic"
)

type SClientUserMgr struct {
	M_SApp *SApp

	M_TCPListener bsn_net.STCPListener

	M_TId2User  map[uint16]*SClientUser
	M_TClientId uint16
	M_SState    bsn_common.SState
	M_MutexUser sync.Mutex
}

func NewSClientUserMgr(vSApp *SApp) (this *SClientUserMgr, err error) {
	GSLog.Debugln("NewSClientUserMgr")
	this = &SClientUserMgr{
		M_SApp:     vSApp,
		M_TId2User: make(map[uint16]*SClientUser),
	}

	this.M_SState.Reset()

	this.M_TCPListener.Reset()
	this.M_TCPListener.SetAddr(fmt.Sprintf(":%v", vSApp.ListenPort()))
	this.M_TCPListener.SetIListenerCB(this)

	return
}

func (this *SClientUserMgr) ListenerCBOnAcceptNetConn(vConn net.Conn) (err error) {
	vClientId := this.genClientId()
	if vClientId == 0 {
		return errors.New("genClientId fail")
	}

	GSLog.Debugf("client connect ClientId=%v \n", vClientId)

	vSClientUser, err := NewSClientUser(this, vConn, vClientId)
	if err != nil {
		return err
	}

	err = this.addClient(vSClientUser)
	if err != nil {
		return err
	}
	vSClientUser.run()
	return nil
}

func (this *SClientUserMgr) ListenerCBOnClose() (err error) {
	GSLog.Mustln("ListenerCBOnClose")
	return nil
}

func (this *SClientUserMgr) start() (err error) {
	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		GSLog.Errorln("can`t start")
		return nil
	}

	defer func() {
		if err != nil {
			this.M_SState.Set(bsn_common.CState_Idle)
			GSLog.Errorln(err)
		}
	}()

	err = this.startTCPListen()
	if err != nil {
		return err
	}

	this.M_SState.Set(bsn_common.CState_Runing)
	return nil
}

func (this *SClientUserMgr) stop() (err error) {
	if !this.M_SState.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		GSLog.Errorln("can`t stop")
		return nil
	}

	defer func() {
		if err != nil {
			this.M_SState.Set(bsn_common.CState_Runing)
			GSLog.Errorln(err)
		}
	}()

	err = this.stopTCPListen()
	if err != nil {
		return err
	}
	this.closeAllClient()

	this.M_SState.Set(bsn_common.CState_Idle)
	return nil
}

func (this *SClientUserMgr) startTCPListen() (err error) {
	return this.M_TCPListener.Start()
}

func (this *SClientUserMgr) stopTCPListen() (err error) {
	return this.M_TCPListener.Stop()
}

// generate clientid
// if not generate return 0
func (this *SClientUserMgr) genClientId() uint16 {
	for i := 0; i < 100; i++ {
		this.M_TClientId++
		if this.M_TClientId == 0 {
			continue
		}

		vSClientUser, _ := this.getClient(this.M_TClientId)
		if vSClientUser == nil {
			return this.M_TClientId
		}
	}
	return 0
}

func (this *SClientUserMgr) getClient(vClientId uint16) (vSClientUser *SClientUser, err error) {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	vSClientUser = this.M_TId2User[vClientId]
	return
}

func (this *SClientUserMgr) addClient(vSClientUser *SClientUser) (err error) {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	if !this.M_SState.Is(bsn_common.CState_Runing) {
		err = errors.New("not run")
		return err
	}

	this.M_TId2User[vSClientUser.ClientId()] = vSClientUser
	return nil
}

func (this *SClientUserMgr) delClient(vTClientId uint16) (err error) {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	delete(this.M_TId2User, vTClientId)
	return nil
}

func (this *SClientUserMgr) closeAllClient() (err error) {
	GSLog.Mustln("closeAllClient")

	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	for _, vClientUser := range this.M_TId2User {
		if vClientUser.setCloseReason(bsn_common.CCloseReason_KickOut) {
			err = vClientUser.close()
			if err != nil {
				GSLog.Errorln(err)
			}
		}
	}

	return
}
