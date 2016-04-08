package bsn_gate_config

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "time""
	// "net"
)

type SServerUserMgr struct {
	*bsn_common.SState

	M_SApp  *SApp
	M_Users []*SServerUser
}

func NewSServerUserMgr(vSApp *SApp) (*SServerUserMgr, error) {
	GSLog.Debugln("NewSServerUserMgr")
	this := &SServerUserMgr{
		M_SApp:  vSApp,
		M_Users: make([]*SServerUser, 1),
	}
	this.SState = bsn_common.NewSState()

	return this, nil
}

func (this *SServerUserMgr) App() *SApp {
	return this.M_SApp
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

	this.Set(bsn_common.CState_Idle)
	return nil
}

func (this *SServerUserMgr) ShowInfo() {
}

func (this *SServerUserMgr) Send(vSClientUser *SClientUser, vSMsgHeader *bsn_msg.SMsgHeader, vbyMsgBody []byte) error {
	GSLog.Debugln("Send")
	GSLog.Mustln(vSClientUser)
	GSLog.Mustln(vSMsgHeader)
	GSLog.Mustln(vbyMsgBody)
	GSLog.Mustln(string(vbyMsgBody))
	return nil
}
