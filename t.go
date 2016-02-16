package main

import (
	// "fmt"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_log"
)

// var _ = mImp1.Main
var GLog *bsn_log.SLog = bsn_log.New()

type SCmd1 struct {
}

func (this *SCmd1) XYZ(strArray []string) {
	GLog.Debugln("SCmd1 XYZ")
}

func (this *SCmd1) XYZ_help(strArray []string) {
	GLog.Debugln("SCmd1 XYZ_help")
}

type SCmd2 struct {
}

func (this *SCmd2) XYZ(strArray []string) {
	GLog.Debugln("SCmd2 XYZ")
}

func (this *SCmd2) XYZ_help(strArray []string) {
	GLog.Debugln("SCmd2 XYZ_help")
}

var GSCmd1 SCmd1
var GSCmd2 SCmd2

func main() {

	GLog.Debugln(1)
	GLog.Debugln(2)
	GLog.Debugln(3)

	bsn_input.GInput.Reg("Log", &GLog.M_SCmd)
	bsn_input.GInput.Reg("Mod1", &GSCmd1)
	bsn_input.GInput.Reg("Mod2", &GSCmd2)
	bsn_input.GInput.Run()
}

/*
gate -id 1 -level 1 -port 40001
gate -id 2 -level 1 -port 40002 -UpNodeAddr localhost:40001
gate -id 3 -level 1 -port 40003 -UpNodeAddr localhost:40001
gate -id 4 -level 1 -port 40004 -UpNodeAddr localhost:40001
*/
