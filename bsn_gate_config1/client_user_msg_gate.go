package bsn_gate3

import (
	"bsn_define"
	"bsn_msg_gate_gateconfig"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (this *SClientUser) procGateMsg(msgType bsn_msg_gate_gateconfig.ECmdGate2GateConfig) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	// case bsn_msg.GMsgDefine_Client2Gate_Ping:
	// 	return this.ProcMsg_Ping()
	// case bsn_msg.GMsgDefine_Client2Gate_Pong:
	// 	return this.ProcMsg_Pong()

	case bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_GetServerConfigReq:
		return this.ProcMsg_CmdGate2GateConfig_GetServerConfigReq()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
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
				VstrAddr:            proto.String("localhost:40001"),
			},
		},
	}
	return this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_GetServerConfigRes), sendMsg)
}
