package bsn_gate3

import (
	"bsn_msg_gate_gateconfig"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_log"
)

func init() {
	GSLog = bsn_log.GSLog
}

type TClientId uint16

func IsMsgGateConfig(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_Max) {
		return false
	}
	return true
}
