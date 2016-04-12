package bsn_2

import (
	"github.com/bsn069/go/bsn_common"
)

type SCmd struct {
	// *SCmdGate
	*SCmdGate2
	*SCmdClient
	*SCmdGateConfig
	*SCmdEcho
}

func NewCmd() *SCmd {
	this := &SCmd{}
	// this.SCmdGate = NewCmdGate()
	this.SCmdGate2 = NewCmdGate2()
	this.SCmdClient = NewCmdClient()
	this.SCmdGateConfig = NewSCmdGateConfig()
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

	this.SCmdEcho.ECHO_RUN(vParams)
	this.SCmdGateConfig.GATE_CONFIG_RUN(vParams)
	this.SCmdGate2.GATE2_RUN(vParams)
}
