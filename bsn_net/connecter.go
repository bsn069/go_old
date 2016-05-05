package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "sync"
)

type INetConnecterImp interface {
	NetConnecterImpRun() error
	NetConnecterImpOnClose() error
}

type SNetConnecter struct {
	*SSessionAddrConnect
	*bsn_common.SState
	*bsn_common.SNotifyClose

	M_INetConnecterImp INetConnecterImp
}

func NewNetConnecter(vINetConnecterImp INetConnecterImp) (*SNetConnecter, error) {
	GSLog.Debugln("NewNetConnecter")

	this := &SNetConnecter{}
	this.M_INetConnecterImp = vINetConnecterImp
	this.SSessionAddrConnect, _ = NewSSessionAddrConnect()
	this.SState = bsn_common.NewSState()
	this.SNotifyClose = bsn_common.NewSNotifyClose()

	return this, nil
}

func (this *SNetConnecter) ShowInfo() {
	GSLog.Mustln("ShowInfo")
}

func (this *SNetConnecter) Run() (err error) {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	this.SNotifyClose.Clear()

	err = this.Connect()
	if err != nil {
		return err
	}

	go this.runImp()
	return nil
}

func (this *SNetConnecter) Close() (err error) {
	if !this.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		return errors.New("not listen")
	}
	GSLog.Debugln("close")

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	}()

	this.Conn().Close()
	this.SNotifyClose.NotifyClose()
	this.SNotifyClose.WaitClose()

	return nil
}

func (this *SNetConnecter) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("close connect")
		this.Conn().Close()
		this.SetConn(nil)

		this.M_INetConnecterImp.NetConnecterImpOnClose()
		this.SNotifyClose.Close()
		this.Set(bsn_common.CState_Idle)
		GSLog.Debugln("close ok")
	}()

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	GSLog.Debugln("run imp")
	vbQuit := false
	for !vbQuit {
		err := this.M_INetConnecterImp.NetConnecterImpRun()
		if err != nil {
			GSLog.Errorln(err)
			vbQuit = true
			break
		}

		select {
		case <-this.SNotifyClose.M_chanNotifyClose:
			GSLog.Mustln("receive a notify to close")
			vbQuit = true
		default:
		}
	}
}
