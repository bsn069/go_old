package bsn_gate3

import (
	// "bsn_define"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"bsn_msg_client_gate"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (this *SClientUser) procClientMsg(msgType bsn_msg_client_gate.ECmdClient2Gate) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_client_gate.ECmdClient2Gate_CmdClient2Gate_TestReq:
		return this.ProcMsg_CmdClient2Gate_TestReq()
	case bsn_msg_client_gate.ECmdClient2Gate_CmdClient2Gate_TestRes:
		return this.ProcMsg_CmdClient2Gate_TestRes()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SClientUser) ProcMsg_CmdClient2Gate_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdClient2Gate_TestReq")

	recvMsg := new(bsn_msg_client_gate.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_client_gate.STestRes{
		VstrInfo: proto.String("gate test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_client_gate.ECmdGate2Client_CmdGate2Client_TestRes), sendMsg)

	return
}

func (this *SClientUser) ProcMsg_CmdClient2Gate_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdClient2Gate_TestRes")

	recvMsg := new(bsn_msg_client_gate.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}
