package bsn_2

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	bsn_app "github.com/bsn069/go/bsn_gate"
	"github.com/bsn069/go/bsn_input"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
)

type SCmdGate struct {
	M_TId2App map[uint32]*bsn_app.SApp
}

func NewCmdGate() *SCmdGate {
	this := &SCmdGate{
		M_TId2App: make(map[uint32]*bsn_app.SApp),
	}
	return this
}

func (this *SCmdGate) createApp(vAppId uint32) (*bsn_app.SApp, error) {
	if _, ok := this.M_TId2App[vAppId]; ok {
		return nil, errors.New("app had exist")
	}

	vSApp, err := bsn_app.NewSApp(vAppId)
	if err != nil {
		return nil, err
	}

	vSApp.M_TFuncAppClose = func() {
		GSLog.Debugln("on app close")
		delete(this.M_TId2App, vAppId)
		bsn_input.GSInput.SetUseMod("Main")
	}

	this.M_TId2App[vAppId] = vSApp
	return vSApp, nil
}

func (this *SCmdGate) GATE(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("appid")
		return
	}

	vuAppId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	_, err = this.createApp(uint32(vuAppId))
	if err != nil {
		GSLog.Errorln(err)
		return
	}
}
