package main

import (
	// "fmt"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_log"
)

// var _ = mImp1.Main
var GSLog *bsn_log.SLog

func init() {
	var vSCmd *bsn_log.SCmd
	GSLog, vSCmd = bsn_log.New()
	bsn_input.GSInput.Reg("Log", bsn_log.GSCmd)
	bsn_input.GSInput.Reg("mainLog", vSCmd)
}

type SCmd1 struct {
}

func (this *SCmd1) XYZ(strArray []string) {
	GSLog.Debugln("SCmd1 XYZ")
}

func (this *SCmd1) XYZ_help(strArray []string) {
	GSLog.Debugln("SCmd1 XYZ_help")
}

type SCmd2 struct {
}

func (this *SCmd2) XYZ(strArray []string) {
	GSLog.Debugln("SCmd2 XYZ")
}

func (this *SCmd2) XYZ_help(strArray []string) {
	GSLog.Debugln("SCmd2 XYZ_help")
}

func main() {
	var vGSCmd1 SCmd1
	var vGSCmd2 SCmd2

	GSLog.Debugln(1)
	GSLog.Debugln(2)
	GSLog.Debugln(3)

	bsn_input.GSInput.Reg("Mod1", &vGSCmd1)
	bsn_input.GSInput.Reg("Mod2", &vGSCmd2)
	bsn_input.GSInput.Run()
}

/*
gate -id 1 -level 1 -port 40001
gate -id 2 -level 1 -port 40002 -UpNodeAddr localhost:40001
gate -id 3 -level 1 -port 40003 -UpNodeAddr localhost:40001
gate -id 4 -level 1 -port 40004 -UpNodeAddr localhost:40001
*/
