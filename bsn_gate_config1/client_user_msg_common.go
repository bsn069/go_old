package bsn_gate3

import (
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) ProcMsg_Ping() error {
	GSLog.Debugln("ProcMsg_Ping")
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Server2Gate_Pong, this.M_by2MsgBody)
	return nil
}

func (this *SClientUser) ProcMsg_Pong() error {
	GSLog.Debugln("ProcMsg_Pong", string(this.M_by2MsgBody))
	return nil
}
