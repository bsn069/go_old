package bsn_echo

import (
	// "bsn_define"
	"bsn_msg_client_echo_server"
	// "github.com/bsn069/go/bsn_common"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
	"github.com/golang/protobuf/proto"
)

func (this *SClientUser) procClientMsg(vMsg *bsn_msg.SMsg_Gate2Server_ClientMsg) error {
	msgType := bsn_msg_client_echo_server.ECmdClient2EchoServer(vMsg.M_SMsgHeader.Type())
	GSLog.Debugln("ClientId= ", vMsg.M_ClientId)
	GSLog.Debugln("msgType= ", msgType)

	switch msgType {
	case bsn_msg_client_echo_server.ECmdClient2EchoServer_CmdClient2EchoServer_TestReq:
		return this.ProcMsg_CmdClient2EchoServer_TestReq(vMsg)
	case bsn_msg_client_echo_server.ECmdClient2EchoServer_CmdClient2EchoServer_TestRes:
		return this.ProcMsg_CmdClient2EchoServer_TestRes(vMsg)
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SClientUser) ProcMsg_CmdClient2EchoServer_TestReq(vMsg *bsn_msg.SMsg_Gate2Server_ClientMsg) (err error) {
	GSLog.Debugln("ProcMsg_CmdClient2EchoServer_TestReq")

	recvMsg := new(bsn_msg_client_echo_server.STestReq)
	if err = proto.Unmarshal(vMsg.M_byMsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(recvMsg.VstrInfo)

	// sendMsg := &bsn_msg_client_echo_server.STestRes{
	// 	VstrInfo: proto.String("echo server test res"),
	// }
	// this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_client_echo_server.ECmdEchoServe2Client_CmdClient2EchoServer_TestRes), sendMsg)

	return
}

func (this *SClientUser) ProcMsg_CmdClient2EchoServer_TestRes(vMsg *bsn_msg.SMsg_Gate2Server_ClientMsg) (err error) {
	GSLog.Debugln("ProcMsg_CmdClient2EchoServer_TestRes")

	recvMsg := new(bsn_msg_client_echo_server.STestRes)
	if err = proto.Unmarshal(vMsg.M_byMsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(recvMsg.VstrInfo)

	return
}
