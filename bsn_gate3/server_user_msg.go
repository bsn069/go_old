package bsn_gate3

import (
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUser) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")

	msgType := this.MsgType()
	switch msgType {
	case bsn_msg.GMsgDefine_Server2Gate_Ping:
		return this.ProcMsg_Ping()
	case bsn_msg.GMsgDefine_Server2Gate_Pong:
		return this.ProcMsg_Pong()
	case bsn_msg.GMsgDefine_Server2Gate_ClientMsg:
		return this.ProcMsg_ClientMsg()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", msgType)
	return errors.New(strInfo)
}
