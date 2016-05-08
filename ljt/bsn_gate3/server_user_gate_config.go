package bsn_gate3

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	"bsn_msg_gate_gateconfig"
	"github.com/golang/protobuf/proto"
	"time"
)

type SServerUserGateConfig struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
	M_Close          bool
}

func NewSServerUserGateConfig(vSServerUserMgr *SServerUserMgr) (*SServerUserGateConfig, error) {
	GSLog.Debugln("NewSServerUserGateConfig()")

	this := &SServerUserGateConfig{
		M_SServerUserMgr: vSServerUserMgr,
		M_Close:          false,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

	return this, nil
}

func (this *SServerUserGateConfig) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUserGateConfig) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	if !this.M_Close {
		GSLog.Debugln("reconnect")
		go func() {
			for {
				time.Sleep(time.Duration(3) * time.Second)
				err := this.Run()
				if err == nil {
					break
				}
				GSLog.Debugln(err)
			}
		}()
	}
	return nil
}

func (this *SServerUserGateConfig) ShowInfo() {
}

func (this *SServerUserGateConfig) send_CmdGate2GateConfig_GetServerConfigReq() error {
	sendMsg := &bsn_msg_gate_gateconfig.SGate2GateConfig_GetServerConfigReq{
		Vu32Id: proto.Uint32(1),
	}
	return this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_GetServerConfigReq), sendMsg)
}
