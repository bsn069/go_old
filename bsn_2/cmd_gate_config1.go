package bsn_2

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	bsn_gate_config "github.com/bsn069/go/bsn_gate_config1"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
)

type SCmdGateConfig1 struct {
	M_SApp *bsn_gate_config.SApp
}

func NewSCmdGateConfig1() *SCmdGateConfig1 {
	this := &SCmdGateConfig1{}
	return this
}

func (this *SCmdGateConfig1) GATE_CONFIG_RUN(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("id")
		return
	}

	if this.M_SApp != nil {
		GSLog.Errorln("this.M_SApp != nil")
		return
	}

	vu32Id, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	this.M_SApp, _ = bsn_gate_config.NewSApp(uint32(vu32Id))

	vuClientListenPort := 51000 + uint16(vu32Id)

	err = this.M_SApp.UserMgr().ClientUserMgr().SetAddr(":" + strconv.Itoa(int(vuClientListenPort)))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	this.M_SApp.Run()
}
