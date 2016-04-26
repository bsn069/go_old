package bsn_echo

import (
	// "bsn_define"
	"bsn_msg_gate_server"
	// "github.com/bsn069/go/bsn_common"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/golang/protobuf/proto"
)

func (this *SClientUser) procGateMsg(msgType bsn_msg_gate_server.ECmdGate2Server) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_ClientMsg:
		return this.ProcMsg_CmdGate2Server_ClientMsg()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SClientUser) ProcMsg_CmdGate2Server_ClientMsg() error {
	vClientMsg := new(bsn_msg.SMsg_Gate2Server_ClientMsg)
	vClientMsg.DeSerialize(this.M_by2MsgBody)
	return this.procClientMsg(vClientMsg)
}
