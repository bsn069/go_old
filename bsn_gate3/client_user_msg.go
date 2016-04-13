package bsn_gate3

import (
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.M_SMsgHeader.Type()
	switch this.M_SMsgHeader.Type() {
	case bsn_msg.GMsgDefine_Gate2Client_Min:
		return this.ProcMsg_GateReg()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", this.M_SMsgHeader.Type())
	return errors.New(strInfo)
}
