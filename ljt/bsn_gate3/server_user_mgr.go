package bsn_gate3

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"time"
	// "net"
	"bsn_msg_gate_gateconfig"
	"strconv"
	// "sync"
)

type SServerUserMgr struct {
	*bsn_common.SState

	M_SUserMgr *SUserMgr
	M_Users    []*SServerUser

	M_SServerUserGateConfig *SServerUserGateConfig

	M_SServerConfigs []*bsn_msg_gate_gateconfig.SServerConfig
}

func NewSServerUserMgr(vSUserMgr *SUserMgr) (*SServerUserMgr, error) {
	GSLog.Debugln("NewSServerUserMgr")
	this := &SServerUserMgr{
		M_SUserMgr: vSUserMgr,
	}
	this.SState = bsn_common.NewSState()

	return this, nil
}

func (this *SServerUserMgr) UserMgr() *SUserMgr {
	return this.M_SUserMgr
}

func (this *SServerUserMgr) Run() (err error) {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	defer func() {
		if err == nil {
			return
		}
		this.M_SServerUserGateConfig.Close()
		this.M_SServerUserGateConfig = nil
	}()

	err = this.run_do_1()
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			return
		}
		this.closeAllUser()
		this.M_Users = nil
	}()

	err = this.run_do_2()
	if err != nil {
		return
	}

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	GSLog.Debugln("run complete")
	return nil
}

func (this *SServerUserMgr) run_do_1() (err error) {
	this.M_SServerUserGateConfig, _ = NewSServerUserGateConfig(this)
	this.M_SServerUserGateConfig.SetAddr("localhost:" + strconv.Itoa(int(bsn_common.GateConfigPort(1))))

	this.M_SServerConfigs = nil

	for {
		err = this.M_SServerUserGateConfig.Run()
		if err == nil {
			break
		}
		GSLog.Errorln(err)
		time.Sleep(time.Duration(5) * time.Second)
	}

	err = this.M_SServerUserGateConfig.send_CmdGate2GateConfig_GetServerConfigReq()
	if err != nil {
		return
	}

	// wait gate config init
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Second)
		if this.M_SServerConfigs != nil {
			return
		}
	}

	return errors.New("gate config init timeout")
}

func (this *SServerUserMgr) run_do_2() (err error) {
	this.M_Users = make([]*SServerUser, len(this.M_SServerConfigs))
	for i, a := range this.M_SServerConfigs {
		GSLog.Debugln(i, a)

		vSServerUser, _ := NewSServerUser(this)
		this.M_Users[i] = vSServerUser
		vSServerUser.SetAddr(a.GetVstrAddr())

		for {
			err = vSServerUser.Run()
			if err == nil {
				break
			}
			GSLog.Errorln(err)
			time.Sleep(time.Duration(5) * time.Second)
		}
	}

	return
}

func (this *SServerUserMgr) closeAllUser() (err error) {
	return this.MapAllUser(func(vSServerUser *SServerUser) {
		vSServerUser.Close()
	})
}

func (this *SServerUserMgr) Close() (err error) {
	if !this.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		return errors.New("not listen")
	}
	GSLog.Debugln("close")

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	}()

	this.M_SServerUserGateConfig.Close()
	this.M_SServerUserGateConfig = nil

	this.closeAllUser()
	this.M_Users = nil

	this.Set(bsn_common.CState_Idle)
	return nil
}

func (this *SServerUserMgr) ShowInfo() {
}

func (this *SServerUserMgr) OnClientMsg(vSClientUser *SClientUser) error {
	return this.MapOneUser(func(vSServerUser *SServerUser) bool {
		return vSServerUser.OnClientMsg(vSClientUser)
	})
}

func (this *SServerUserMgr) Ping(strInfo string) error {
	return this.MapAllUser(func(vSServerUser *SServerUser) {
		vSServerUser.Ping([]byte(strInfo))
	})
}

func (this *SServerUserMgr) MapAllUser(mapFunc func(vSServerUser *SServerUser)) error {
	for _, vServerUser := range this.M_Users {
		if vServerUser == nil {
			continue
		}
		mapFunc(vServerUser)
	}
	return nil
}

func (this *SServerUserMgr) MapOneUser(mapFunc func(vSServerUser *SServerUser) bool) error {
	for _, vServerUser := range this.M_Users {
		if vServerUser == nil {
			continue
		}
		if mapFunc(vServerUser) {
			return nil
		}
	}
	return nil
}
