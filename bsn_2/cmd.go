package bsn_2

import (
	"github.com/bsn069/go/bsn_common"
)

type SCmd struct {
	*SCmdGate
	*SCmdClient
}

func NewCmd() *SCmd {
	this := &SCmd{}
	this.SCmdGate = NewCmdGate()
	this.SCmdClient = NewCmdClient()

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
