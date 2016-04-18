package bsn_gate3

import (
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()
	switch msgType {
	case bsn_msg.GMsgDefine_Client2Gate_Ping:
		return this.ProcMsg_Ping()
	case bsn_msg.GMsgDefine_Client2Gate_Pong:
		return this.ProcMsg_Pong()

	case bsn_msg.GMsgDefine_Gate2GateConfig_Reg:
		return this.ProcMsg_Reg()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", msgType)
	return errors.New(strInfo)
}
