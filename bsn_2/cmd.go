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
