package bsn_gate

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
	M_SGate *SGate
}

func (this *SCmd) LS(vTInputParams bsn_common.TInputParams) {

}

func (this *SCmd) LS_help(vTInputParams bsn_common.TInputParams) {
	GSLog.Mustln("    list all mod")
}
