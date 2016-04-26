package bsn_gate3

import (
	"bsn_define"
	"errors"
	"fmt"
	// "bsn_msg_gate_gateconfig"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/golang/protobuf/proto"
)

func (this *SServerUserGateConfig) procSysMsg(msgType bsn_define.ECmd) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	// case bsn_msg.GMsgDefine_Client2Gate_Ping:
	// 	return this.ProcMsg_Ping()
	// case bsn_msg.GMsgDefine_Client2Gate_Pong:
	// 	return this.ProcMsg_Pong()

	case bsn_define.ECmd_Cmd_Ping:
		return this.ProcMsg_Cmd_Ping()
	case bsn_define.ECmd_Cmd_Pong:
		return this.ProcMsg_Cmd_Pong()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUserGateConfig) ProcMsg_Cmd_Ping() (err error) {
	GSLog.Debugln("ProcMsg_Cmd_Ping")

	return this.Pong(this.M_by2MsgBody)
}

func (this *SServerUserGateConfig) ProcMsg_Cmd_Pong() (err error) {
	GSLog.Debugln("ProcMsg_Cmd_Pong ", string(this.M_by2MsgBody))
	return nil
}
