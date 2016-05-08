package bsn_gate3

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	"bsn_define"
	"bsn_msg_gate_server"
	"github.com/golang/protobuf/proto"
)

func (this *SServerUser) send_CmdGate2Server_ClientMsg(vSClientUser *SClientUser) bool {
	vSMsg_Gate2Server_ClientMsg := new(bsn_msg.SMsg_Gate2Server_ClientMsg)
	vSMsg_Gate2Server_ClientMsg.Fill(uint16(vSClientUser.Id()), vSClientUser.M_SMsgHeader, vSClientUser.M_by2MsgBody)
	this.SendMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_ClientMsg), vSMsg_Gate2Server_ClientMsg.Serialize())

	return true
}

func (this *SServerUser) send_CmdGate2Server_LoginReq(vEServerType bsn_define.EServerType) (err error) {
	vSApp := this.UserMgr().UserMgr().App()

	sendMsg := &bsn_msg_gate_server.SLoginReq{
		Id:         proto.Uint32(vSApp.Id()),
		ServerType: vEServerType.Enum(),
	}

	return this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_LoginReq), sendMsg)
}
