package bsn_2

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	app "github.com/bsn069/go/bsn_echo"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
)

type SCmdEcho struct {
	M_SApp *app.SApp
}

func NewSCmdEcho() *SCmdEcho {
	this := &SCmdEcho{}
	return this
}

func (this *SCmdEcho) ECHO_RUN(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("id")
		return
	}

	if this.M_SApp != nil {
		GSLog.Errorln("this.M_SApp != nil")
		return
	}

	vu32Id, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	this.M_SApp, _ = app.NewSApp(uint32(vu32Id))

	vuClientListenPort := 52000 + uint16(vu32Id)
	err = this.M_SApp.GetClientMgr().SetAddr(":" + strconv.Itoa(int(vuClientListenPort)))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	this.M_SApp.Run()
}
