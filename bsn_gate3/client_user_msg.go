package bsn_gate3

import (
	// "errors"
	// "fmt"
	"bsn_define"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	return this.UserMgr().UserMgr().ServerUserMgr().OnClientMsg(this)
}
