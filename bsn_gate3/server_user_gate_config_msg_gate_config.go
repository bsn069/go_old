package bsn_gate3

import (
	// "bsn_define"
	"bsn_msg_gate_gateconfig"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/golang/protobuf/proto"
)

func (this *SServerUserGateConfig) procGateConfigMsg(msgType bsn_msg_gate_gateconfig.ECmdGateConfig2Gate) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_TestReq:
		return this.ProcMsg_CmdGateConfig2Gate_TestReq()
	case bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_TestRes:
		return this.ProcMsg_CmdGateConfig2Gate_TestRes()

	case bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_GetServerConfigRes:
		return this.ProcMsg_CmdGateConfig2Gate_GetServerConfigRes()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUserGateConfig) ProcMsg_CmdGateConfig2Gate_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdGateConfig2Gate_TestReq")

	recvMsg := new(bsn_msg_gate_gateconfig.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_gate_gateconfig.STestRes{
		VstrInfo: proto.String("gate test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_TestRes), sendMsg)

	return
}

func (this *SServerUserGateConfig) ProcMsg_CmdGateConfig2Gate_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdGateConfig2Gate_TestRes")

	recvMsg := new(bsn_msg_gate_gateconfig.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}

func (this *SServerUserGateConfig) ProcMsg_CmdGateConfig2Gate_GetServerConfigRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdGateConfig2Gate_GetServerConfigRes")

	recvMsg := new(bsn_msg_gate_gateconfig.SGateConfig2Gate_GetServerConfigRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}

	this.UserMgr().M_SServerConfigs = recvMsg.VSServerConfigs
	this.UserMgr().M_chanWaitGateConfig <- true

	return
}
