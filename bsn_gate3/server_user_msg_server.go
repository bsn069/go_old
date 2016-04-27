package bsn_gate3

import (
	"bsn_msg_gate_server"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/golang/protobuf/proto"
)

func (this *SServerUser) procServerMsg(msgType bsn_msg_gate_server.ECmdServe2Gate) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_TestReq:
		return this.ProcMsg_CmdServer2Gate_TestReq()
	case bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_TestRes:
		return this.ProcMsg_CmdServer2Gate_TestRes()

	case bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_ClientMsg:
		return this.ProcMsg_CmdServer2Gate_ClientMsg()
	case bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_GetClientMsgRangeRes:
		return this.ProcMsg_CmdServer2Gate_GetClientMsgRangeRes()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUser) ProcMsg_CmdServer2Gate_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdServer2Gate_TestReq")

	recvMsg := new(bsn_msg_gate_server.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_gate_server.STestRes{
		VstrInfo: proto.String("gate test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_TestRes), sendMsg)

	return
}

func (this *SServerUser) ProcMsg_CmdServer2Gate_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdServer2Gate_TestRes")

	recvMsg := new(bsn_msg_gate_server.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}

func (this *SServerUser) ProcMsg_CmdServer2Gate_ClientMsg() error {
	GSLog.Debugln("ProcMsg_CmdServer2Gate_ClientMsg")
	vSMsg_Server2Gate_ClientMsg := new(bsn_msg.SMsg_Server2Gate_ClientMsg)
	vSMsg_Server2Gate_ClientMsg.DeSerialize(this.M_by2MsgBody)
	return this.UserMgr().UserMgr().ClientUserMgr().Send(TClientId(vSMsg_Server2Gate_ClientMsg.M_ClientId), vSMsg_Server2Gate_ClientMsg.M_byMsg)
}

func (this *SServerUser) ProcMsg_CmdServer2Gate_GetClientMsgRangeRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdServer2Gate_GetClientMsgRangeRes")

	recvMsg := new(bsn_msg_gate_server.SGetClientMsgRangeRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(recvMsg.GetVu32MsgTypeMin(), recvMsg.GetVu32MsgTypeMax())

	this.M_MsgTypeMin = bsn_common.TMsgType(recvMsg.GetVu32MsgTypeMin())
	this.M_MsgTypeMax = bsn_common.TMsgType(recvMsg.GetVu32MsgTypeMax())

	return
}
