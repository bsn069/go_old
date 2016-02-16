package bsn_input

import (
	"github.com/bsn069/go/bsn_log"
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
	*bsn_log.SCmd
	M_u16QuitCode uint16
}

func (this *SCmd) LS(strArray []string) {
	for _, vModName := range GInput.M_TUpperName2RegName {
		GLog.Mustln(vModName)
	}
}

func (this *SCmd) LS_help(vTParams TParams) {
	GLog.Mustln("    list all mod")
}

func (this *SCmd) QUIT(strArray []string) {
	if this.M_u16QuitCode == 0 || len(strArray) != 1 {
		this.M_u16QuitCode = uint16(rand.Uint32())
		if this.M_u16QuitCode < 10000 {
			this.M_u16QuitCode += 10000
		}
		GLog.Mustf("input [quit %d]\n", this.M_u16QuitCode)
		return
	}
	if strconv.Itoa(int(this.M_u16QuitCode)) != strArray[0] {
		this.M_u16QuitCode = 0
		GLog.Mustln("quit cancel")
		return
	}

	GInput.Quit()
}

func (this *SCmd) QUIT_help(vTParams TParams) {
	GLog.Mustln("    quit input")
}

func (this *SCmd) CD(strArray []string) {
	switch strArray[0] {
	case "..", "/":
		GInput.M_strUseMod = ""
	default:
		err := GInput.SetUseMod(strArray[0])
		if err != nil {
			GLog.Errorln(err)
		}
	}
}

func (this *SCmd) CD_help(vTParams TParams) {
	GLog.Mustln("    enter mod")
}
