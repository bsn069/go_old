package bsn_gate3

import (
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) ProcMsg_Reg() error {
	GSLog.Debugln("ProcMsg_Reg")
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_GateConfig2Gate_Reg, nil)
	return nil
}
