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
	"strconv"
	// "strings"
)

type SCmd struct {
	M_SGate *SGate
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

func (this *SCmd) SERVER_LISTEN_ADDR(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a string like localhost:40000")
		return
	}

	err := this.M_SGate.GetServerMgr().SetListenAddr(vTInputParams[0])
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) CLIENT_LISTEN(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetClientMgr().Listen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) SERVER_LISTEN(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetServerMgr().Listen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) CLIENT_LISTEN_STOP(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetClientMgr().StopListen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) SERVER_LISTEN_STOP(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetServerMgr().StopListen()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) LISTEN(vTInputParams bsn_common.TInputParams) {
	this.M_SGate.Listen()
}

func (this *SCmd) STOP_LISTEN(vTInputParams bsn_common.TInputParams) {
	this.M_SGate.Listen()
}

func (this *SCmd) CLOSE(vTInputParams bsn_common.TInputParams) {
	this.M_SGate.Close()
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

func (this *SCmd) SERVER_LISTEN_PORT(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("ListenPort")
		return
	}

	vuListenPort, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = this.M_SGate.GetServerMgr().SetListenPort(uint16(vuListenPort))
	if err != nil {
		GSLog.Errorln(err)
		return
	}
}

func (this *SCmd) CLIENT_CLOSE(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetClientMgr().Close()
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) SERVER_CLOSE(vTInputParams bsn_common.TInputParams) {
	err := this.M_SGate.GetServerMgr().Close()
	if err != nil {
		GSLog.Errorln(err)
	}
}
