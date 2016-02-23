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

func (this *SCmd) SET_CLIENT_LISTEN_ADDR(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a string like localhost:40000")
		return
	}

	err := this.M_SGate.GetClientMgr().SetListenAddr(vTInputParams[0])
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) SET_SERVER_LISTEN_ADDR(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a string like localhost:40000")
		return
	}

	err := this.M_SGate.GetServerMgr().SetListenAddr(vTInputParams[0])
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) LISTEN_CLIENT(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetClientMgr().Listen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) LISTEN_SERVER(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetServerMgr().Listen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) STOP_LISTEN_CLIENT(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetClientMgr().StopListen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) STOP_LISTEN_SERVER(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetServerMgr().StopListen()
	if err != nil {
		GSLog.Errorln(err)
	}
}
