package bsn_2

import (
	"github.com/bsn069/go/bsn_common"
	"strconv"
)

type SCmd struct {
	M_SCmdGate *SCmdGate
}

func NewCmd() *SCmd {
	this := &SCmd{}
	this.M_SCmdGate = NewCmdGate()

	return this
}

func (this *SCmd) TEST(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a number")
		return
	}
}

func (this *SCmd) TEST_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    SETOUTMASK_help")
}

func (this *SCmd) GATE_TEST(vTInputParams bsn_common.TInputParams) {
	this.M_SCmdGate.Test()
}

func (this *SCmd) GATE_CREATE(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a number with gate id")
		return
	}

	vuGateId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	_, err = this.M_SCmdGate.CreateGate(bsn_common.TGateId(vuGateId))
	if err != nil {
		GSLog.Errorln(err)
		return
	}
}

func (this *SCmd) GATE_CREATE_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    create gate")
}

func (this *SCmd) GATE(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 3 {
		GSLog.Errorln("gateid clientListenPort serverListenPort")
		return
	}

	vuGateId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vuClientListenPort, err := strconv.ParseUint(vTInputParams[1], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vuServerListenPort, err := strconv.ParseUint(vTInputParams[2], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSGate, err := this.M_SCmdGate.CreateGate(bsn_common.TGateId(vuGateId))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = vSGate.GetServerMgr().SetListenPort(uint16(vuServerListenPort))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = vSGate.GetClientMgr().SetListenPort(uint16(vuClientListenPort))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSGate.Listen()
}

func (this *SCmd) GATE_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    gate")
}
