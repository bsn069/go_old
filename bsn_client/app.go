package bsn_client

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	"github.com/bsn069/go/bsn_input"
	"strconv"
	// "sync"
	// "time"
)

type TAppFuncOnQuit func()
type SApp struct {
	*bsn_common.SRunCmd
	M_Id   uint32
	M_Name string

	M_HadStart       bool
	M_TAppFuncOnQuit TAppFuncOnQuit

	M_SGate *SGate
}

func NewSApp(vu32Id uint32) (this *SApp, err error) {
	GSLog.Debugln("NewSApp() vu32Id=", vu32Id)
	this = &SApp{
		M_Id: vu32Id,
	}
	this.M_Name = GAppName + strconv.Itoa(int(this.Id()))
	this.SRunCmd, _ = bsn_common.NewSRunCmd(this.runCmdFuncOnQuit, GSLog)
	this.M_SGate, _ = NewSGate(this)

	vSCmd, err := NewSCmd(this)
	if err != nil {
		return nil, err
	}

	GSLog.Mustln("reg all cmd")
	this.RegCmd("start", this.start)
	this.RegCmd("stop", this.stop)

	bsn_input.GSInput.Reg(this.Name(), vSCmd)
	return this, nil
}

func (this *SApp) Id() uint32 {
	return this.M_Id
}

func (this *SApp) Name() string {
	return this.M_Name
}

func (this *SApp) ListenPort() uint16 {
	return bsn_common.GatePort(this.Id())
}

func (this *SApp) runCmdFuncOnQuit() (err error) {
	err = bsn_input.GSInput.UnReg(this.Name())
	if err != nil {
		panic(err)
	}

	if this.M_HadStart {
		err = this.stop()
		if err != nil {
			panic(err)
		}
	}

	if this.M_TAppFuncOnQuit != nil {
		this.M_TAppFuncOnQuit()
		this.M_TAppFuncOnQuit = nil
	}
	return
}

func (this *SApp) start() (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if this.M_HadStart {
		GSLog.Errorln("had start")
		return
	}

	err = this.M_SGate.start()
	if err != nil {
		return
	}

	GSLog.Debugln("start")
	this.M_HadStart = true
	return
}

func (this *SApp) stop() (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_HadStart {
		GSLog.Errorln("not start")
		return
	}

	err = this.M_SGate.stop()
	if err != nil {
		return
	}

	GSLog.Debugln("stop")
	this.M_HadStart = false
	return
}
