package bsn_2

import (
	"github.com/bsn069/go/bsn_common"
	"strconv"
)

type SCmd struct {
	// *SCmdGate
	// *SCmdGate2
	*SCmdGate3
	*SCmdGate
	*SCmdClient1
	*SCmdGateConfig1
	*SCmdEcho
}

func NewCmd() *SCmd {
	this := &SCmd{}
	// this.SCmdGate = NewCmdGate()
	// this.SCmdGate2 = NewCmdGate2()
	this.SCmdGate = NewCmdGate()
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
	if len(vTInputParams) != 1 {
		GSLog.Errorln("appid")
		return
	}
	this.ECHO_RUN(vTInputParams)
	this.GATE_CONFIG_RUN(vTInputParams)
	this.GATE3_RUN(vTInputParams)
	this.CLIENT1_RUN(vTInputParams)
}

func (this *SCmd) TEST1(vTInputParams bsn_common.TInputParams) {
	vParam := make(bsn_common.TInputParams, 1)

	vParam[0] = "1"
	this.ECHO_RUN(vParam)
	this.GATE_CONFIG_RUN(vParam)

	for i := 1; i <= 1; i++ {
		vParam[0] = strconv.Itoa(i)
		this.GATE3_RUN(vParam)
	}
}
