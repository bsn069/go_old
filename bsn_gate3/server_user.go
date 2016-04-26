package bsn_gate3

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	"bsn_msg_gate_server"
)

type SServerUser struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
	M_ServerType     uint8

	M_MsgTypeMin bsn_common.TMsgType
	M_MsgTypeMax bsn_common.TMsgType
}

func NewSServerUser(vSServerUserMgr *SServerUserMgr) (*SServerUser, error) {
	GSLog.Debugln("NewSServerUser()")

	this := &SServerUser{
		M_SServerUserMgr: vSServerUserMgr,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

	return this, nil
}

func (this *SServerUser) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUser) ServerType() uint8 {
	return this.M_ServerType
}

func (this *SServerUser) OnClientMsg(vSClientUser *SClientUser) bool {
	if vSClientUser.MsgType() < this.M_MsgTypeMin || vSClientUser.MsgType() > this.M_MsgTypeMax {
		return false
	}

	vSMsg_Gate2Server_ClientMsg := new(bsn_msg.SMsg_Gate2Server_ClientMsg)
	vSMsg_Gate2Server_ClientMsg.Fill(uint16(vSClientUser.Id()), vSClientUser.M_SMsgHeader, vSClientUser.M_by2MsgBody)
	this.SendMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_ClientMsg), vSMsg_Gate2Server_ClientMsg.Serialize())

	return true
}

func (this *SServerUser) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerUser) ShowInfo() {
}
