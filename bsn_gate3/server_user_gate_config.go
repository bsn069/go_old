package bsn_gate3

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
)

type SServerUserGateConfig struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
}

func NewSServerUserGateConfig(vSServerUserMgr *SServerUserMgr) (*SServerUserGateConfig, error) {
	GSLog.Debugln("NewSServerUserGateConfig()")

	this := &SServerUserGateConfig{
		M_SServerUserMgr: vSServerUserMgr,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

	return this, nil
}

func (this *SServerUserGateConfig) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUserGateConfig) Run() {
	this.SConnecterWithMsgHeader.Run()
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2GateConfig_Reg, nil)
}

func (this *SServerUserGateConfig) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerUserGateConfig) ShowInfo() {
}

func (this *SServerUserGateConfig) Ping(strInfo string) error {
	return this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2Server_Ping, []byte(strInfo))
}
