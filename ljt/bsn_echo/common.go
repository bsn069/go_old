package bsn_echo

import (
	"bsn_msg_gate_server"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_log"
)

func init() {
	GSLog = bsn_log.GSLog
}

type TClientId uint16

func IsMsgGate(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_Max) {
		return false
	}
	return true
}
