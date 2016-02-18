package bsn_input

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_log"
)

func init() {
	var vSCmd *bsn_log.SCmd
	GSLog, vSCmd = bsn_log.New()
	GSInput = &SInput{
		M_TInputUpperName2RegName:   make(bsn_common.TInputUpperName2RegName),
		M_TInputUpperName2CmdStruct: make(bsn_common.TInputUpperName2CmdStruct),
		M_SCmd: new(SCmd),
	}

	GSInput.Reg("InputLog", vSCmd)
}
