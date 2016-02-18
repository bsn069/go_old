/*
Package bsn_input.

type SCmd1 struct {
}

func (this *SCmd1) XYZ(strArray []string) {
	GSLog.Debugln("SCmd1 XYZ")
}

func (this *SCmd1) XYZ_help(strArray []string) {
	GSLog.Debugln("SCmd1 XYZ_help")
}

var GSCmd1 SCmd1
bsn_input.GSInput.Reg("Mod1", &GSCmd1)
bsn_input.GSInput.Run()
*/
package bsn_input

import (
	"github.com/bsn069/go/bsn_log"
)

var GSLog *bsn_log.SLog
var GSInput *SInput
