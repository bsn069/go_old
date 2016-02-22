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
	GSLog = bsn_log.GSLog
}

type SCmd1 struct {
}

func (this *SCmd1) XYZ(strArray []string) {
	GSLog.Debugln("SCmd1 XYZ")
}

func (this *SCmd1) XYZ_help(strArray []string) {
	GSLog.Debugln("SCmd1 XYZ_help")
}

func main() {
	var vGSCmd1 SCmd1

	GSLog.Debugln(1)
	GSLog.Debugln(2)
	GSLog.Debugln(3)

	bsn_input.GSInput.Reg("Mod1", &vGSCmd1)
	bsn_input.GSInput.Run()
}

/*
gate -id 1 -level 1 -port 40001
gate -id 2 -level 1 -port 40002 -UpNodeAddr localhost:40001
gate -id 3 -level 1 -port 40003 -UpNodeAddr localhost:40001
gate -id 4 -level 1 -port 40004 -UpNodeAddr localhost:40001
*/
