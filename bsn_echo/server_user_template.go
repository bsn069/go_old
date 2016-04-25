package bsn_echo

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	// "bsn_msg_gate_gateconfig"
	// "github.com/golang/protobuf/proto"
)

type SServerUserTemplate struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
}

func NewSServerUserTemplate(vSServerUserMgr *SServerUserMgr) (*SServerUserTemplate, error) {
	GSLog.Debugln("NewSServerUserTemplate()")

	this := &SServerUserTemplate{
		M_SServerUserMgr: vSServerUserMgr,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

	return this, nil
}

func (this *SServerUserTemplate) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUserTemplate) Run() {
	this.SConnecterWithMsgHeader.Run()

	// sendMsg := &bsn_msg_gate_gateconfig.SGate2GateConfig_GetServerConfigReq{
	// 	Vu32Id: proto.Uint32(1),
	// }
	// this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_GetServerConfigReq), sendMsg)
}

func (this *SServerUserTemplate) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerUserTemplate) ShowInfo() {
}
