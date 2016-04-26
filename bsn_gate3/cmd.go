package bsn_gate3

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
	"strings"
)

type SCmd struct {
	M_SApp *SApp
}

func (this *SCmd) SHOW_INFO(vTInputParams bsn_common.TInputParams) {
	this.M_SApp.ShowInfo()
}

func (this *SCmd) CLIENT_LISTEN_ADDR(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("param must a string like localhost:40000")
		return
	}

	err := this.M_SApp.UserMgr().ClientUserMgr().SetAddr(vTInputParams[0])
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) CLOSE(vTInputParams bsn_common.TInputParams) {
	this.M_SApp.Close()
}

func (this *SCmd) RUN(vTInputParams bsn_common.TInputParams) {
	this.M_SApp.Run()
}

func (this *SCmd) SERVERS_PING(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) < 1 {
		GSLog.Errorln("msg strings")
		return
	}

	strMsg := strings.Join([]string(vTInputParams), " ")

	err := this.M_SApp.UserMgr().ServerUserMgr().Ping(strMsg)
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) CLIENTS_PING(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) < 1 {
		GSLog.Errorln("msg strings")
		return
	}

	strMsg := strings.Join([]string(vTInputParams), " ")

	err := this.M_SApp.UserMgr().ClientUserMgr().Ping(strMsg)
	if err != nil {
		GSLog.Errorln(err)
	}
}

func (this *SCmd) CLIENTS_TEST_REQ(vTInputParams bsn_common.TInputParams) {
	err := this.M_SApp.UserMgr().ClientUserMgr().TestReq()
	if err != nil {
		GSLog.Errorln(err)
	}
}
