package bsn_gate3

import (
	// "errors"
	// "fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUserGateConfig) ProcMsg_Ping() error {
	GSLog.Debugln("ProcMsg_Ping")
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2Server_Pong, nil)
	return nil
}

func (this *SServerUserGateConfig) ProcMsg_Pong() error {
	GSLog.Debugln("ProcMsg_Pong", string(this.M_by2MsgBody))
	return nil
}
