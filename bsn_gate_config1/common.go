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

func IsMsgGate(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_gate_gateconfig.ECmdGate2GateConfig_CmdGate2GateConfig_Max) {
		return false
	}
	return true
}
