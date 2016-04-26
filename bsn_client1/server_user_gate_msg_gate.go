package bsn_client1

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"bsn_msg_client_gate"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (this *SServerUserGate) procGateMsg(msgType bsn_msg_client_gate.ECmdGate2Client) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_client_gate.ECmdGate2Client_CmdGate2Client_TestReq:
		return this.ProcMsg_CmdGate2Client_TestReq()
	case bsn_msg_client_gate.ECmdGate2Client_CmdGate2Client_TestRes:
		return this.ProcMsg_CmdGate2Client_TestRes()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUserGate) ProcMsg_CmdGate2Client_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2Client_TestReq")

	recvMsg := new(bsn_msg_client_gate.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_client_gate.STestRes{
		VstrInfo: proto.String("client test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_client_gate.ECmdClient2Gate_CmdClient2Gate_TestRes), sendMsg)

	return
}

func (this *SServerUserGate) ProcMsg_CmdGate2Client_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2Client_TestRes")

	recvMsg := new(bsn_msg_client_gate.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}
