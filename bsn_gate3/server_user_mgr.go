package bsn_gate3

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
	M_Users    []*SServerUser

	M_SServerUserGateConfig *SServerUserGateConfig
	M_chanWaitGateConfig    chan bool
}

func NewSServerUserMgr(vSUserMgr *SUserMgr) (*SServerUserMgr, error) {
	GSLog.Debugln("NewSServerUserMgr")
	this := &SServerUserMgr{
		M_SUserMgr:           vSUserMgr,
		M_Users:              make([]*SServerUser, 1),
		M_chanWaitGateConfig: make(chan bool, 1),
	}
	this.SState = bsn_common.NewSState()
	this.M_SServerUserGateConfig, _ = NewSServerUserGateConfig(this)
	this.M_SServerUserGateConfig.SetAddr("localhost:30001")

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

	this.M_SServerUserGateConfig.Run()
	<-this.M_chanWaitGateConfig

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

func (this *SServerUserMgr) OnClientMsg(vSClientUser *SClientUser) error {
	for _, vServerUser := range this.M_Users {
		if vServerUser == nil {
			continue
		}
		if vServerUser.OnClientMsg(vSClientUser) {
			GSLog.Debugln(vServerUser.ServerType(), "proc msg")
			return nil
		}
	}

	return errors.New("unknown msg")
}

func (this *SServerUserMgr) Ping(strInfo string) error {
	for _, vServerUser := range this.M_Users {
		if vServerUser == nil {
			continue
		}
		vServerUser.Ping(strInfo)
	}
	return nil
}
