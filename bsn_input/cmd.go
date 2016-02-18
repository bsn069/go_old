package bsn_input

import (
	"github.com/bsn069/go/bsn_common"
	// "errors"
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	"math/rand"
	// "reflect"
	"strconv"
	// "strings"
)

type SCmd struct {
	M_u16QuitCode uint16
}

func (this *SCmd) LS(vTInputParams bsn_common.TInputParams) {
	for _, vModName := range GSInput.M_TInputUpperName2RegName {
		GSLog.Mustln(vModName)
	}
}

func (this *SCmd) LS_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    list all mod")
}

func (this *SCmd) QUIT(vTInputParams bsn_common.TInputParams) {
	if this.M_u16QuitCode == 0 || len(vTInputParams) != 1 {
		this.M_u16QuitCode = uint16(rand.Uint32())
		if this.M_u16QuitCode < 10000 {
			this.M_u16QuitCode += 10000
		}
		GSLog.Mustf("input [quit %d]\n", this.M_u16QuitCode)
		return
	}
	if strconv.Itoa(int(this.M_u16QuitCode)) != vTInputParams[0] {
		this.M_u16QuitCode = 0
		GSLog.Mustln("quit cancel")
		return
	}

	GSInput.Quit()
}

func (this *SCmd) QUIT_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    quit input")
}

func (this *SCmd) CD(vTInputParams bsn_common.TInputParams) {
	switch vTInputParams[0] {
	case "..", "/":
		GSInput.M_strUseMod = ""
	default:
		err := GSInput.SetUseMod(vTInputParams[0])
		if err != nil {
			GSLog.Errorln(err)
		}
	}
}

func (this *SCmd) CD_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    enter mod")
}
