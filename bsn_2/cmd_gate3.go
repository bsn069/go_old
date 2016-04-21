package bsn_2

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	bsn_app "github.com/bsn069/go/bsn_gate3"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
)

type SCmdGate3 struct {
	M_TId2App map[uint32]*bsn_app.SApp
	M_AppId   uint32
}

func NewCmdGate3() *SCmdGate3 {
	this := &SCmdGate3{
		M_TId2App: make(map[uint32]*bsn_app.SApp),
		M_AppId:   0,
	}
	return this
}

func (this *SCmdGate3) GATE3_RUN(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("appid")
		return
	}

	vuAppId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSApp, err := this.Gate3Create(uint32(vuAppId))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vuClientListenPort := 20000 + vuAppId
	err = vSApp.UserMgr().ClientUserMgr().SetAddr(":" + strconv.Itoa(int(vuClientListenPort)))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSApp.Run()
}

func (this *SCmdGate3) Gate3Create(vAppId uint32) (*bsn_app.SApp, error) {
	vSApp, err := bsn_app.NewSApp(vAppId)
	if err != nil {
		return nil, err
	}

	this.M_TId2App[vAppId] = vSApp
	return vSApp, nil
}
