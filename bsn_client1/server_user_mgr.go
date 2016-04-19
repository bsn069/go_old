package bsn_client1

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "time""
	// "net"
)

type SServerUserMgr struct {
	*bsn_common.SState

	M_SUserMgr *SUserMgr

	M_SServerUserGate *SServerUserGate
}

func NewSServerUserMgr(vSUserMgr *SUserMgr) (*SServerUserMgr, error) {
	GSLog.Debugln("NewSServerUserMgr")
	this := &SServerUserMgr{
		M_SUserMgr: vSUserMgr,
	}
	this.SState = bsn_common.NewSState()
	this.M_SServerUserGate, _ = NewSServerUserGate(this)
	this.M_SServerUserGate.SetAddr("localhost:40001")

	return this, nil
}

func (this *SServerUserMgr) UserMgr() *SUserMgr {
	return this.M_SUserMgr
}

func (this *SServerUserMgr) Run() (err error) {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	this.M_SServerUserGate.Run()
	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)

	return nil
}

func (this *SServerUserMgr) Close() (err error) {
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

	this.M_SServerUserGate.Close()
	this.Set(bsn_common.CState_Idle)

	return nil
}

func (this *SServerUserMgr) ShowInfo() {
	this.M_SServerUserGate.ShowInfo()
}
