/*
Package bsn_log.

*/
package bsn_2

import (
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_log"
	// "runtime"
)

var GSLog *bsn_log.SLog
var GSCmd *SCmd

func Run() {
	// runtime.GOMAXPROCS(runtime.NumCPU()*2 + 2)
	bsn_input.GSInput.Run()
	GSLog.Mustln("quit")
}
