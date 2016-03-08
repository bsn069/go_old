package bsn_2

import (
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_log"
)

func init() {
	GSLog = bsn_log.GSLog
	GSCmd = NewCmd()

	bsn_input.GSInput.Reg("Main", GSCmd)
	bsn_input.GSInput.SetUseMod("Main")
}
