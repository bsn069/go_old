package bsn_input

import (
	"github.com/bsn069/go/bsn_common"
	// "errors"
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

type SCmd struct {
	m_quitCode  uint16
	m_bShowHelp bool
}

func (this *SCmd) LS(strArray []string) {
	if this.m_bShowHelp {
		GLog.Mustln("    list all mod")
		return
	}

	for vModName, _ := range gInput.m_tMod2Chan {
		GLog.Mustln(vModName)
	}
}

func (this *SCmd) QUIT(strArray []string) {
	if this.m_bShowHelp {
		GLog.Mustln("    quit input")
		return
	}

	if this.m_quitCode == 0 || len(strArray) != 1 {
		this.m_quitCode = uint16(rand.Uint32())
		if this.m_quitCode < 10000 {
			this.m_quitCode += 10000
		}
		GLog.Mustf("input [quit %d]\n", this.m_quitCode)
		return
	}
	if strconv.Itoa(int(this.m_quitCode)) != strArray[0] {
		this.m_quitCode = 0
		GLog.Mustln("quit cancel")
		return
	}

	instance().quit()
}

func (this *SCmd) CD(strArray []string) {
	if len(strArray) != 1 {
		this.m_bShowHelp = true
	}
	if this.m_bShowHelp {
		GLog.Mustln("    enter mod")
		return
	}

	switch strArray[0] {
	case "..", "/":
		instance().m_strUseMod = ""
	default:
		instance().setUseMod(strArray[0])
	}
}

func (this *SCmd) HELP(strArray []string) {
	if this.m_bShowHelp {
		GLog.Mustln("    show help info")
		return
	}

	this.m_bShowHelp = true
	if len(strArray) > 0 {
		err := bsn_common.CallStructFunc(this, strArray[0], strArray)
		if err == nil {
			this.m_bShowHelp = false
		}
	}
	if this.m_bShowHelp {
		t := reflect.TypeOf(this)
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			lowerName := strings.ToLower(method.Name)
			GLog.Mustln(lowerName)
			bsn_common.CallStructFunc(this, method.Name, strArray)
		}
	}
}
