package bsn_client1

import (
	"bsn_msg_client_echo_server"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_log"
)

func init() {
	GSLog = bsn_log.GSLog
}

type TClientId uint16

func IsMsgEchoServer(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_client_echo_server.ECmdEchoServe2Client_CmdEchoServer2Client_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_client_echo_server.ECmdEchoServe2Client_CmdEchoServer2Client_Max) {
		return false
	}
	return true
}
