package bsn_gate3

import (
	"bsn_define"
	"bsn_msg_gate_gateconfig"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	if IsMsgGate(msgType) {
		return this.procGateMsg(bsn_msg_gate_gateconfig.ECmdGate2GateConfig(msgType))
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func IsMsgGate(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_Max) {
		return false
	}
	return true
}
