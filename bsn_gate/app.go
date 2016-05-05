package bsn_gate

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_log"
	"github.com/bsn069/go/bsn_input"
	"strconv"
	"sync"
)

type TFuncAppClose func()
type SApp struct {
	M_Id             uint32
	M_Name           string
	M_SServerUserMgr *SServerUserMgr
	M_HadRun         bool
	M_HadClose       bool
	M_Mutex          *sync.Mutex
	M_TFuncAppClose  TFuncAppClose
}

func NewSApp(vu32Id uint32) (this *SApp, err error) {
	GSLog.Debugln("NewSApp() vu32Id=", vu32Id)
	this = &SApp{
		M_Id:       vu32Id,
		M_Mutex:    new(sync.Mutex),
		M_HadClose: false,
		M_HadRun:   false,
	}
	this.M_Name = GAppName + strconv.Itoa(int(this.Id()))

	vSCmd, err := NewSCmd(this)
	if err != nil {
		return nil, err
	}

	this.M_SServerUserMgr, err = NewSServerUserMgr(this)
	if err != nil {
		return nil, err
	}

	bsn_input.GSInput.Reg(this.Name(), vSCmd)
	return this, nil
}

func (this *SApp) Id() uint32 {
	return this.M_Id
}

func (this *SApp) Name() string {
	return this.M_Name
}

func (this *SApp) ServerUserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SApp) Run() {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()

	if this.M_HadRun {
		GSLog.Errorln("had run, only run once")
		return
	}

	err := this.M_SServerUserMgr.run()
	if err != nil {
		panic(err)
	}

	this.M_HadRun = true
}

func (this *SApp) Close() {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()

	if this.M_HadClose {
		GSLog.Errorln("had close, only close once")
		return
	}

	err := bsn_input.GSInput.UnReg(this.Name())
	if err != nil {
		panic(err)
	}

	if this.M_HadRun {
		err = this.M_SServerUserMgr.close()
		if err != nil {
			panic(err)
		}
	}

	this.M_HadClose = true
	if this.M_TFuncAppClose != nil {
		this.M_TFuncAppClose()
		this.M_TFuncAppClose = nil
	}
}

func (this *SApp) ListenPort() uint16 {
	return bsn_common.GatePort(this.Id())
}
