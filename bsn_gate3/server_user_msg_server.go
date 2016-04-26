package bsn_gate3

import (
	"bsn_msg_gate_server"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SServerUser) procServerMsg(msgType bsn_msg_gate_server.ECmdServe2Gate) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_ClientMsg:
		return this.ProcMsg_CmdServer2Gate_ClientMsg()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUser) ProcMsg_CmdServer2Gate_ClientMsg() error {
	GSLog.Debugln("ProcMsg_CmdServer2Gate_ClientMsg")
	vSMsg_Server2Gate_ClientMsg := new(bsn_msg.SMsg_Server2Gate_ClientMsg)
	vSMsg_Server2Gate_ClientMsg.DeSerialize(this.M_by2MsgBody)
	return this.UserMgr().UserMgr().ClientUserMgr().Send(TClientId(vSMsg_Server2Gate_ClientMsg.M_ClientId), vSMsg_Server2Gate_ClientMsg.M_byMsg)
}
