package bsn_gate2

import (
	"github.com/bsn069/go/bsn_common"
	// "errors"
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	// "math/rand"
	// "reflect"
	"strconv"
	// "strings"
)

type SCmd struct {
	M_SGate *SGate
}

func (this *SCmd) SHOW_INFO(vTInputParams bsn_common.TInputParams) {
	this.M_SGate.ShowInfo()
}

func (this *SCmd) CLIENT_LISTEN_ADDR(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a string like localhost:40000")
		return
	}

	err := this.M_SGate.GetClientMgr().SetListenAddr(vTInputParams[0])
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) CLOSE(vTInputParams bsn_common.TInputParams) {
	this.M_SGate.Close()
}

func (this *SCmd) RUN(vTInputParams bsn_common.TInputParams) {
	this.M_SGate.Run()
}

func (this *SCmd) CLIENT_LISTEN_PORT(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("ListenPort")
		return
	}

	vuListenPort, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = this.M_SGate.GetClientMgr().SetListenPort(uint16(vuListenPort))
	if err != nil {
		GSLog.Errorln(err)
		return
	}
}
