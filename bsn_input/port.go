/*
Package bsn_input.

type SCmd1 struct {
}

func (this *SCmd1) XYZ(strArray []string) {
	GLog.Debugln("SCmd1 XYZ")
}

func (this *SCmd1) XYZ_help(strArray []string) {
	GLog.Debugln("SCmd1 XYZ_help")
}

var GSCmd1 SCmd1
bsn_input.GInput.Reg("Mod1", &GSCmd1)
bsn_input.GInput.Run()
*/
package bsn_input

import (
	"github.com/bsn069/go/bsn_log"
)

var GLog *bsn_log.SLog
var GInput *SInput

func init() {
	GLog = bsn_log.New()
	GInput = &SInput{
		M_TUpperName2RegName:   make(TUpperName2RegName),
		M_TUpperName2CmdStruct: make(TUpperName2CmdStruct),
		M_SCmd:                 new(SCmd),
	}

	GInput.M_SCmd.SCmd = &GLog.M_SCmd
}
