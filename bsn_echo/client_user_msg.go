package bsn_gate3

import (
	"bsn_define"
	"bsn_msg_gate_server"
	"errors"
	"fmt"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	if IsMsgServer(msgType) {
		return this.procServerMsg(bsn_msg_gate_server.ECmdGate2Server(msgType))
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}
