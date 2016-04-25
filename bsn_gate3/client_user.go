package bsn_gate3

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "unsafe"
	// "net"
	// "sync"
)

type SClientUser struct {
	*bsn_net.SSessionWithMsgHeader
	*bsn_common.SState
	*bsn_common.SNotifyClose

	M_SClientUserMgr *SClientUserMgr

	M_TClientId TClientId
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
	}
	this.SSessionWithMsgHeader, _ = bsn_net.NewSSessionWithMsgHeader()
	this.SState = bsn_common.NewSState()
	this.SNotifyClose = bsn_common.NewSNotifyClose()
	return this, nil
}

func (this *SClientUser) UserMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SClientUser) SetId(vTClientId TClientId) error {
	this.M_TClientId = vTClientId
	return nil
}

func (this *SClientUser) Id() TClientId {
	return this.M_TClientId
}

func (this *SClientUser) Run() (err error) {
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

	go this.runImp()
	return nil
}

func (this *SClientUser) Close() (err error) {
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
	// this.SNotifyClose.NotifyClose()
	this.SNotifyClose.WaitClose()

	return nil
}

func (this *SClientUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("close connect")
		this.Conn().Close()

		GSLog.Debugln("delete from user mgr")
		this.M_SClientUserMgr.delClient(this.Id())

		this.SNotifyClose.Close()
		this.Set(bsn_common.CState_Idle)
		GSLog.Debugln("close ok")
	}()

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	for {
		err := this.RecvMsg()
		if err != nil {
			GSLog.Errorln(err)
			break
		}

		err = this.procMsg()
		if err != nil {
			GSLog.Errorln(err)
			break
		}
	}
}
