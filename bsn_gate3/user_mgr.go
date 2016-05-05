package bsn_gate3

import (
// "errors"
// "github.com/bsn069/go/bsn_common"
// "github.com/bsn069/go/bsn_msg"
// "time""
// "net"
)

type SUserMgr struct {
	M_SApp           *SApp
	M_SClientUserMgr *SClientUserMgr
	M_SServerUserMgr *SServerUserMgr
}

func NewSUserMgr(vSApp *SApp) (this *SUserMgr, err error) {
	GSLog.Debugln("NewSUserMgr")
	this = &SUserMgr{
		M_SApp: vSApp,
	}

	this.M_SClientUserMgr, err = NewSClientUserMgr(this)
	if err != nil {
		GSLog.Errorln("NewSClientUserMgr fail")
		return nil, err
	}

	this.M_SServerUserMgr, err = NewSServerUserMgr(this)
	if err != nil {
		GSLog.Errorln("NewSServerUserMgr fail")
		return nil, err
	}

	return this, nil
}

func (this *SUserMgr) App() *SApp {
	return this.M_SApp
}

func (this *SUserMgr) ClientUserMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SUserMgr) ServerUserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SUserMgr) ShowInfo() {
	this.ClientUserMgr().ShowInfo()
	this.ServerUserMgr().ShowInfo()
}

func (this *SUserMgr) Close() {
	this.ClientUserMgr().Close()
	this.ServerUserMgr().Close()
}

func (this *SUserMgr) Run() {
	err := this.ServerUserMgr().Run()
	if err != nil {
		panic(err)
	}

	err = this.ClientUserMgr().Run()
	if err != nil {
		panic(err)
	}
}
