package bsn_client

import (
	"github.com/bsn069/go/bsn_common"
	// "errors"
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	// "math/rand"
	// "reflect"
	// "strconv"
	// "strings"
)

type SCmd struct {
	M_SApp *SApp
}

func NewSCmd(vSApp *SApp) (this *SCmd, err error) {
	GSLog.Debugln("NewSCmd()")

	this = &SCmd{
		M_SApp: vSApp,
	}

	return this, nil
}

func (this *SCmd) DO(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("need func name")
		return
	}

	this.M_SApp.Do(vTInputParams[0])
}
