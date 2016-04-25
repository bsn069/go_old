package bsn_echo

import (
	"bsn_define"
	// "bsn_msg_gate_gateconfig"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUserTemplate) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	// if IsMsgServer(msgType) {
	// 	return this.procServerMsg(bsn_msg_gate_gateconfig.ECmdGateConfig2Gate(msgType))
	// }

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}
