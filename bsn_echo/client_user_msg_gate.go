package bsn_gate3

import (
	// "bsn_define"
	"bsn_msg_gate_server"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"errors"
	"fmt"
	// "github.com/golang/protobuf/proto"
)

func (this *SClientUser) procGateMsg(msgType bsn_msg_gate_server.ECmdGate2Server) error {
	GSLog.Debugln("msgType=", msgType)

	// switch msgType {
	// // case bsn_msg.GMsgDefine_Client2Gate_Ping:
	// // 	return this.ProcMsg_Ping()
	// // case bsn_msg.GMsgDefine_Client2Gate_Pong:
	// // 	return this.ProcMsg_Pong()

	// case bsn_msg_gate_server.ECmdGate2GateConfig_CmdGate2GateConfig_GetServerConfigReq:
	// 	return this.ProcMsg_CmdGate2GateConfig_GetServerConfigReq()
	// }

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}
