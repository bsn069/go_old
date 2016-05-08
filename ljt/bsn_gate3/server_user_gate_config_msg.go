package bsn_gate3

import (
	"bsn_define"
	"bsn_msg_gate_gateconfig"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUserGateConfig) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	if IsMsgGateConfig(msgType) {
		return this.procGateConfigMsg(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate(msgType))
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func IsMsgGateConfig(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_Max) {
		return false
	}
	return true
}
