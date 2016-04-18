package bsn_gate3

import (
	// "errors"
	// "fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUser) ProcMsg_ClientMsg() error {
	GSLog.Debugln("ProcMsg_ClientMsg")
	vSMsg_Server2Gate_ClientMsg := new(bsn_msg.SMsg_Server2Gate_ClientMsg)
	vSMsg_Server2Gate_ClientMsg.DeSerialize(this.M_by2MsgBody)
	return this.UserMgr().UserMgr().ClientUserMgr().Send(TClientId(vSMsg_Server2Gate_ClientMsg.M_ClientId), vSMsg_Server2Gate_ClientMsg.M_byMsg)
}
