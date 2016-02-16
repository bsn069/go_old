package bsn_gate

import (
	"github.com/bsn069/go/bsn_input"
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
}

func (this *SCmd) Show(strArray []string) {
	if this.m_bShowHelp {
		GLog.Mustln("    list all mod")
		return
	}

	for vModName, _ := range gInput.m_tMod2Chan {
		GLog.Mustln(vModName)
	}
}
