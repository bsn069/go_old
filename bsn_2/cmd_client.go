package bsn_2

import (
	"errors"
	bsn_app "github.com/bsn069/go/bsn_client"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
)

type SCmdClient struct {
	M_TId2App map[uint32]*bsn_app.SApp
}

func NewSCmdClient() *SCmdClient {
	this := &SCmdClient{
		M_TId2App: make(map[uint32]*bsn_app.SApp),
	}
	return this
}

func (this *SCmdClient) CLIENT_DO(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("do func")
		return
	}

	for _, vApp := range this.M_TId2App {
		vApp.Do(vTInputParams[0])
	}
}

func (this *SCmdClient) createApp(vAppId uint32) (*bsn_app.SApp, error) {
	if _, ok := this.M_TId2App[vAppId]; ok {
		return nil, errors.New("app had exist")
	}

	vSApp, err := bsn_app.NewSApp(vAppId)
	if err != nil {
		return nil, err
	}

	vSApp.M_TAppFuncOnQuit = func() {
		GSLog.Debugln("on app close")
		delete(this.M_TId2App, vAppId)
		bsn_input.GSInput.SetUseMod("Main")
	}

	this.M_TId2App[vAppId] = vSApp
	return vSApp, nil
}

func (this *SCmdClient) CLIENT(vTInputParams bsn_common.TInputParams) {
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

func (this *SCmdClient) CLIENT_CREATE(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("connect num")
		return
	}

	vNum, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	if vNum > 1000 {
		vNum = 1000
	}

	vCount := len(this.M_TId2App)
	for i := 0; i < int(vNum); i++ {
		vTInputParams[0] = strconv.Itoa(i + vCount + 1)
		this.CLIENT(vTInputParams)
	}
}

func (this *SCmdClient) CLIENT_TEST1(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("connect num")
		return
	}

	vNum, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	if vNum > 1000 {
		vNum = 1000
	}

	for i := 0; i < int(vNum); i++ {
		vTInputParams[0] = "start"
		this.CLIENT_DO(vTInputParams)
		vTInputParams[0] = "stop"
		this.CLIENT_DO(vTInputParams)
	}
}
