package bsn_client1

import (
	"bsn_define"
	"bsn_msg_client_echo_server"
	"errors"
	"fmt"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUserGate) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	if IsMsgEchoServer(msgType) {
		return this.procEchoServerMsg(bsn_msg_client_echo_server.ECmdEchoServe2Client(msgType))
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}
