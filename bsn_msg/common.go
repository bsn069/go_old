package bsn_msg

import (
	"bsn_define"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_log"
	// "unsafe"
)

func init() {
	GSLog = bsn_log.GSLog
}

func IsMsgSys(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_define.ECmd_Cmd_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_define.ECmd_Cmd_Max) {
		return false
	}
	return true
}
