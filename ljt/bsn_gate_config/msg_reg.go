package bsn_gate_config

import (
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	switch this.M_SMsgHeader.Type() {
	case bsn_msg.GMsgDefine_Gate2GateConfig_GateReg:
		return this.ProcMsg_GateReg()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", this.M_SMsgHeader.Type())
	return errors.New(strInfo)
}
