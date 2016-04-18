package bsn_2

import (
	"github.com/bsn069/go/bsn_common"
)

type SCmd struct {
	// *SCmdGate
	// *SCmdGate2
	*SCmdGate3
	*SCmdClient1
	*SCmdGateConfig1
	*SCmdEcho
}

func NewCmd() *SCmd {
	this := &SCmd{}
	// this.SCmdGate = NewCmdGate()
	// this.SCmdGate2 = NewCmdGate2()
	this.SCmdGate3 = NewCmdGate3()
	this.SCmdClient1 = NewCmdClient1()
	this.SCmdGateConfig1 = NewSCmdGateConfig1()
	this.SCmdEcho = NewSCmdEcho()

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

func (this *SCmd) RUN(vTInputParams bsn_common.TInputParams) {
	vParams := bsn_common.MakeInputParam("1")

	// this.SCmdEcho.ECHO_RUN(vParams)
	this.SCmdGateConfig1.GATE_CONFIG_RUN(vParams)
	this.SCmdGate3.GATE3_RUN(vParams)
	this.SCmdClient1.CLIENT1_RUN(vParams)
}
