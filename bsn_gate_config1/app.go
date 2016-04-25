package bsn_gate3

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_log"
	"strconv"
)

type SApp struct {
	M_SUserMgr *SUserMgr
	M_Id       uint32
}

func NewSApp(vId uint32) (this *SApp, err error) {
	GSLog.Debugln("NewSApp() vId=", vId)
	this = &SApp{
		M_Id: vId,
	}

	this.M_SUserMgr, err = NewSUserMgr(this)
	if err != nil {
		GSLog.Errorln("NewSUserMgr fail")
		return nil, err
	}

	vSCmd := &SCmd{M_SApp: this}
	bsn_input.GSInput.Reg(GAppName+strconv.Itoa(int(vId)), vSCmd)

	return this, nil
}

func (this *SApp) Id() uint32 {
	return this.M_Id
}

func (this *SApp) ConfigListenPort() uint16 {
	return bsn_common.GateConfigPort(this.Id())
}

func (this *SApp) UserMgr() *SUserMgr {
	return this.M_SUserMgr
}

func (this *SApp) Run() {
	this.M_SUserMgr.Run()
}

func (this *SApp) Close() {
	this.M_SUserMgr.Close()
}

func (this *SApp) ShowInfo() {
	GSLog.Mustln("GAppName=", GAppName)
	this.M_SUserMgr.ShowInfo()
}
