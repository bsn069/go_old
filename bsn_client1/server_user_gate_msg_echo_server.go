package bsn_client1

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"bsn_msg_client_echo_server"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (this *SServerUserGate) procEchoServerMsg(msgType bsn_msg_client_echo_server.ECmdEchoServe2Client) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_client_echo_server.ECmdEchoServe2Client_CmdEchoServer2Client_TestReq:
		return this.ProcMsg_CmdEchoServer2Client_TestReq()
	case bsn_msg_client_echo_server.ECmdEchoServe2Client_CmdEchoServer2Client_TestRes:
		return this.ProcMsg_CmdEchoServer2Client_TestRes()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUserGate) ProcMsg_CmdEchoServer2Client_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdEchoServer2Client_TestReq")

	recvMsg := new(bsn_msg_client_echo_server.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_client_echo_server.STestRes{
		VstrInfo: proto.String("client test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_client_echo_server.ECmdClient2EchoServer_CmdClient2EchoServer_TestRes), sendMsg)

	return
}

func (this *SServerUserGate) ProcMsg_CmdEchoServer2Client_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdEchoServer2Client_TestRes")

	recvMsg := new(bsn_msg_client_echo_server.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}
