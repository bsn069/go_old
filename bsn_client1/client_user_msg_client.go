package bsn_client1

import (
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) ProcMsg_Ping() error {
	GSLog.Debugln("ProcMsg_Ping")
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2Client_Pong, nil)
	return nil
}

func (this *SClientUser) ProcMsg_Echo() error {
	GSLog.Debugln("ProcMsg_Echo")
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2Client_Echo, this.M_by2MsgBody)
	return nil
}
