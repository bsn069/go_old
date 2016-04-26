package bsn_gate3

import (
	"bsn_msg_gate_server"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUser) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	return this.procServerMsg(bsn_msg_gate_server.ECmdServe2Gate(msgType))
}
