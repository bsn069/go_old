package bsn_gate3

import (
// "errors"
// "fmt"
// "github.com/bsn069/go/bsn_msg"
)

func (this *SServerUserGateConfig) ProcMsg_Reg() error {
	GSLog.Debugln("ProcMsg_Reg")
	this.UserMgr().M_chanWaitGateConfig <- true
	return nil
}
