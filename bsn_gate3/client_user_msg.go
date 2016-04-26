package bsn_gate3

import (
	// "errors"
	// "fmt"
	"bsn_define"
	"bsn_msg_client_gate"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	msgType := this.MsgType()

	if bsn_msg.IsMsgSys(msgType) {
		return this.procSysMsg(bsn_define.ECmd(msgType))
	}

	if IsMsgClient(msgType) {
		return this.procClientMsg(bsn_msg_client_gate.ECmdClient2Gate(msgType))
	}

	return this.UserMgr().UserMgr().ServerUserMgr().OnClientMsg(this)
}

func IsMsgClient(vTMsgType bsn_common.TMsgType) bool {
	if vTMsgType < bsn_common.TMsgType(bsn_msg_client_gate.ECmdClient2Gate_CmdClient2Gate_Min) {
		return false
	}
	if vTMsgType > bsn_common.TMsgType(bsn_msg_client_gate.ECmdClient2Gate_CmdClient2Gate_Max) {
		return false
	}
	return true
}
