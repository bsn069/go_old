package bsn_gate3

import (
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) ProcMsg_GateReg() error {
	GSLog.Debugln("ProcMsg_GateReg")
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_GateConfig2Gate_GateReg, nil)

	return nil
}
