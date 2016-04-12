package bsn_echo

import (
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_log"
	"strconv"
)

type SApp struct {
	M_SClientUserMgr *SClientUserMgr
	M_SServerUserMgr *SServerUserMgr
	M_Id             uint32
}

func NewSApp(vId uint32) (this *SApp, err error) {
	GSLog.Debugln("NewSApp() vId=", vId)
	this = &SApp{
		M_Id: vId,
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

	vSCmd := &SCmd{M_SApp: this}
	bsn_input.GSInput.Reg(GAppName+strconv.Itoa(int(vId)), vSCmd)

	return this, nil
}

func (this *SApp) ShowInfo() {
	GSLog.Mustln("GAppName=", GAppName)
	this.GetClientMgr().ShowInfo()
	this.GetServerMgr().ShowInfo()
}

func (this *SApp) GetClientMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SApp) GetServerMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SApp) Close() {
	this.GetClientMgr().Close()
	this.GetServerMgr().Close()
}

func (this *SApp) Run() {
	this.GetServerMgr().Run()
	this.GetClientMgr().Run()
}
