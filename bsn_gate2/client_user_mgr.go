package bsn_gate2

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	"net"
	"sync"
)

type TId2ClientUser map[TClientId]*SClientUser
type SClientUserMgr struct {
	M_SGate *SGate

	M_MutexUser sync.Mutex
	M_TId2User  TId2ClientUser
	M_TClientId TClientId

	*bsn_net.SListen

	M_Mutex           sync.Mutex
	M_chanWaitClose   chan bool
	M_chanNotifyClose chan bool
	M_bRun            bool
}

func NewSClientUserMgr(vSGate *SGate) (*SClientUserMgr, error) {
	GSLog.Debugln("NewSClientUserMgr")
	this := &SClientUserMgr{
		M_SGate:           vSGate,
		M_TClientId:       0,
		M_TId2User:        make(TId2ClientUser, 100),
		M_bRun:            false,
		M_chanNotifyClose: make(chan bool, 1),
		M_chanWaitClose:   make(chan bool, 1),
	}
	this.SListen = bsn_net.NewListen()

	return this, nil
}

func (this *SClientUserMgr) ShowInfo() {
	GSLog.Mustln("listen addr: ", this.M_strAddr)
	GSLog.Mustln("max user id: ", this.M_TClientId)
	GSLog.Mustln("user count : ", len(this.M_TId2User))
	GSLog.Mustln("is running : ", this.M_bRun)
	GSLog.Mustln("is listen  : ", this.IsListen())
}

func (this *SClientUserMgr) Send(vTClientId TClientId, vbyMsg []byte) error {
	return nil
}

func (this *SClientUserMgr) Run() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if this.M_bRun {
		return errors.New("running")
	}

	this.Listen()
	go this.runImp()
	this.M_bRun = true
	return nil
}

func (this *SClientUserMgr) Close() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if !this.M_bRun {
		return errors.New("not running")
	}
	GSLog.Mustln("Close begin")

	this.StopListen()
	// clear close chanel
	select {
	case <-this.M_chanNotifyClose:
	default:
	}
	this.M_chanNotifyClose <- true
	// wait close complete
	<-this.M_chanWaitClose

	GSLog.Mustln("Close end")
	return nil
}

func (this *SClientUserMgr) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("on closing")
		this.M_bRun = false

		GSLog.Debugln("close listen")
		this.StopListen()

		GSLog.Debugln("close all not process connect")
		var vbQuit bool = false
		for !vbQuit {
			select {
			case conn := <-this.SListen.M_chanConn:
				conn.Close()
			default:
				vbQuit = true
			}
		}

		GSLog.Debugln("close all user")
		vDelUsers := make([]*SClientUser, len(this.M_TId2User))
		i := 0
		for _, vSClientUser := range this.M_TId2User {
			vDelUsers[i] = vSClientUser
		}
		for _, vSClientUser := range vDelUsers {
			vSClientUser.Close()
		}
		GSLog.Debugln("len(this.M_TId2User)=", len(this.M_TId2User))

		GSLog.Debugln("send notify to wait close chan")
		select {
		case <-this.M_chanWaitClose:
		default:
		}
		this.M_chanWaitClose <- true

		GSLog.Debugln("close all ok")
	}()

	GSLog.Debugln("run imp")
	var vbQuit bool = false
	for !vbQuit {
		vbQuit = true
		select {
		case vConn, ok := <-this.SListen.M_chanConn:
			if !ok {
				GSLog.Errorln("!ok")
				break
			}
			err := this.onNewUser(vConn)
			if err != nil {
				GSLog.Errorln(err)
				vConn.Close()
				break
			}

			vbQuit = false
		case <-this.M_chanNotifyClose:
			GSLog.Mustln("receive a notify to close")
		}
	}
}

func (this *SClientUserMgr) onNewUser(vConn net.Conn) error {
	vSClientUser, err := NewSClientUser(this)
	if err != nil {
		return err
	}

	vTClientId := this.genClientId()
	if vTClientId == 0 {
		return errors.New("genClientId fail")
	}

	err = vSClientUser.SetConn(vConn)
	if err != nil {
		return err
	}

	vSClientUser.SetId(vTClientId)
	this.addClient(vSClientUser)

	vSClientUser.Run()
	return nil
}

func (this *SClientUserMgr) addClient(vSClientUser *SClientUser) error {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	this.M_TId2User[vSClientUser.Id()] = vSClientUser
	return nil
}

func (this *SClientUserMgr) delClient(vTClientId TClientId) error {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	delete(this.M_TId2User, vTClientId)
	return nil
}

func (this *SClientUserMgr) Gate() *SGate {
	return this.M_SGate
}
func (this *SClientUserMgr) getClient(vTClientId TClientId) *SClientUser {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	return this.M_TId2User[vTClientId]
}

// generate clientid
// if not generate return 0
func (this *SClientUserMgr) genClientId() TClientId {
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
