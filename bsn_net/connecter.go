package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"sync"
)

type INetConnecterImp interface {
	NetConnecterImpRun() error
	NetConnecterImpOnClose() error
}

type SNetConnecter struct {
	SSessionAddrConnect
	M_bRun             bool
	M_Mutex            sync.Mutex
	M_chanNotifyClose  chan bool
	M_chanWaitClose    chan bool
	M_INetConnecterImp INetConnecterImp
}

func NewNetConnecter() (*SNetConnecter, error) {
	GSLog.Debugln("NewNetConnecter")

	this := &SNetConnecter{
		M_chanNotifyClose: make(chan bool, 0),
		M_chanWaitClose:   make(chan bool, 0),
	}

	return this, nil
}

func (this *SNetConnecter) Uninit() {
	this.SetConn(nil)
	this.M_INetConnecterImp = nil
}

func (this *SNetConnecter) ShowInfo() {
	GSLog.Mustln("ShowInfo")
}

func (this *SNetConnecter) Run() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if this.M_bRun {
		return errors.New("running")
	}

	err := this.Connect()
	if err != nil {
		return err
	}

	go this.runImp()
	this.M_bRun = true
	return nil
}

func (this *SNetConnecter) Close() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if !this.M_bRun {
		return errors.New("not running")
	}
	GSLog.Mustln("Close begin")

	this.Conn().Close()
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

func (this *SNetConnecter) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("on closing")
		this.M_bRun = false

		GSLog.Debugln("close connect")
		this.Conn().Close()

		GSLog.Debugln("send notify to wait close chan")
		select {
		case <-this.M_chanWaitClose:
		default:
		}
		this.M_chanWaitClose <- true

		this.M_INetConnecterImp.NetConnecterImpOnClose()

		GSLog.Debugln("close ok")
	}()

	for {
		err := this.M_INetConnecterImp.NetConnecterImpRun()
		if err != nil {
			GSLog.Errorln(err)
			break
		}
	}
}
