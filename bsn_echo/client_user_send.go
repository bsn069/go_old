package bsn_echo

import (
	"github.com/bsn069/go/bsn_common"
	// "unsafe"
	// "net"
	// "sync"
	"bsn_msg_client_echo_server"
	"bsn_msg_gate_server"
	"github.com/golang/protobuf/proto"
)

func (this *SClientUser) send_CmdServer2Gate_GetClientMsgRangeRes() error {
	sendMsg := &bsn_msg_gate_server.SGetClientMsgRangeRes{
		Vu32MsgTypeMin: proto.Uint32(uint32(bsn_msg_client_echo_server.ECmdClient2EchoServer_CmdClient2EchoServer_Min)),
		Vu32MsgTypeMax: proto.Uint32(uint32(bsn_msg_client_echo_server.ECmdClient2EchoServer_CmdClient2EchoServer_Max)),
	}
	return this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_GetClientMsgRangeRes), sendMsg)
}

func (this *SClientUser) send_CmdServer2Gate_LoginRes(vResult bsn_msg_gate_server.SLoginRes_EResult) error {
	sendMsg := &bsn_msg_gate_server.SLoginRes{
		Result: vResult.Enum(),
	}
	return this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_LoginRes), sendMsg)
}
