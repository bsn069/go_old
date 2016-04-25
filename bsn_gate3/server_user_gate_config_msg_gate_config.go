package bsn_gate3

import (
	// "bsn_define"
	"bsn_msg_gate_gateconfig"
	"errors"
	"fmt"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/golang/protobuf/proto"
)

func (this *SServerUserGateConfig) procGateConfigMsg(msgType bsn_msg_gate_gateconfig.ECmdGateConfig2Gate) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	// case bsn_msg.GMsgDefine_Client2Gate_Ping:
	// 	return this.ProcMsg_Ping()
	// case bsn_msg.GMsgDefine_Client2Gate_Pong:
	// 	return this.ProcMsg_Pong()

	case bsn_msg_gate_gateconfig.ECmdGateConfig2Gate_CmdGateConfig2Gate_GetServerConfigRes:
		return this.ProcMsg_CmdGateConfig2Gate_GetServerConfigRes()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUserGateConfig) ProcMsg_CmdGateConfig2Gate_GetServerConfigRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdGateConfig2Gate_GetServerConfigRes")

	recvMsg := new(bsn_msg_gate_gateconfig.SGateConfig2Gate_GetServerConfigRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}

	this.UserMgr().M_SServerConfigs = recvMsg.VSServerConfigs
	this.UserMgr().M_chanWaitGateConfig <- true

	return
}
