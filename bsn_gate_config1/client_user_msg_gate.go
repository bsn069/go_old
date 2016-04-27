package bsn_gate3

import (
	"bsn_define"
	"bsn_msg_gate_gateconfig"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
)

func (this *SClientUser) procGateMsg(msgType bsn_msg_gate_gateconfig.ECmdGate2GateConfig) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_TestReq:
		return this.ProcMsg_CmdGate2GateConfig_TestReq()
	case bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_TestRes:
		return this.ProcMsg_CmdGate2GateConfig_TestRes()

	case bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_GetServerConfigReq:
		return this.ProcMsg_CmdGate2GateConfig_GetServerConfigReq()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SClientUser) ProcMsg_CmdGate2GateConfig_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2GateConfig_TestReq")

	recvMsg := new(bsn_msg_gate_gateconfig.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_gate_gateconfig.STestRes{
		VstrInfo: proto.String("gate config test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_TestRes), sendMsg)

	return
}

func (this *SClientUser) ProcMsg_CmdGate2GateConfig_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2GateConfig_TestRes")

	recvMsg := new(bsn_msg_gate_gateconfig.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}

func (this *SClientUser) ProcMsg_CmdGate2GateConfig_GetServerConfigReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2GateConfig_GetServerConfigReq")

	recvMsg := new(bsn_msg_gate_gateconfig.SGate2GateConfig_GetServerConfigReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}

	sendMsg := &bsn_msg_gate_gateconfig.SGateConfig2Gate_GetServerConfigRes{
		VSServerConfigs: []*bsn_msg_gate_gateconfig.SServerConfig{
			&bsn_msg_gate_gateconfig.SServerConfig{
				Vcommon_EServerType: bsn_define.EServerType_ServerType_Echo.Enum(),
				VstrAddr:            proto.String("localhost:" + strconv.Itoa(int(bsn_common.ServerPort(uint32(bsn_define.EServerType_ServerType_Echo), this.UserMgr().UserMgr().App().Id())))),
			},
		},
	}
	return this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_GetServerConfigRes), sendMsg)
}
