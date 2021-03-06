package bsn_echo

import (
	// "bsn_define"
	// "bsn_msg_client_echo_server"
	"bsn_msg_gate_server"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/golang/protobuf/proto"
)

func (this *SClientUser) procGateMsg(msgType bsn_msg_gate_server.ECmdGate2Server) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_TestReq:
		return this.ProcMsg_CmdGate2Server_TestReq()
	case bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_TestRes:
		return this.ProcMsg_CmdGate2Server_TestRes()
	case bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_LoginReq:
		return this.ProcMsg_CmdGate2Server_LoginReq()

	case bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_ClientMsg:
		return this.ProcMsg_CmdGate2Server_ClientMsg()
	case bsn_msg_gate_server.ECmdGate2Server_CmdGate2Server_GetClientMsgRangeReq:
		return this.ProcMsg_CmdGate2Server_GetClientMsgRangeReq()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SClientUser) ProcMsg_CmdGate2Server_TestReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2Server_TestReq")

	recvMsg := new(bsn_msg_gate_server.STestReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	sendMsg := &bsn_msg_gate_server.STestRes{
		VstrInfo: proto.String("echo test res"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_TestRes), sendMsg)

	return
}

func (this *SClientUser) ProcMsg_CmdGate2Server_TestRes() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2Server_TestRes")

	recvMsg := new(bsn_msg_gate_server.STestRes)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	GSLog.Debugln(*recvMsg.VstrInfo)

	return
}

func (this *SClientUser) ProcMsg_CmdGate2Server_ClientMsg() error {
	vClientMsg := new(bsn_msg.SMsg_Gate2Server_ClientMsg)
	vClientMsg.DeSerialize(this.M_by2MsgBody)
	return this.procClientMsg(vClientMsg)
}

func (this *SClientUser) ProcMsg_CmdGate2Server_GetClientMsgRangeReq() error {
	return this.send_CmdServer2Gate_GetClientMsgRangeRes()
}

func (this *SClientUser) ProcMsg_CmdGate2Server_LoginReq() (err error) {
	GSLog.Debugln("ProcMsg_CmdGate2Server_LoginReq")

	recvMsg := new(bsn_msg_gate_server.SLoginReq)
	if err = proto.Unmarshal(this.M_by2MsgBody, recvMsg); err != nil {
		return
	}
	vTClientId := TClientId(recvMsg.GetId())
	GSLog.Debugln("vTClientId= ", vTClientId)

	if vTClientId == 0 {
		err = errors.New("error clientid")
		return
	}

	vResult := bsn_msg_gate_server.SLoginRes_Success

	oldClient := this.UserMgr().getClient(vTClientId)
	if oldClient != nil {
		if oldClient == this {
			vResult = bsn_msg_gate_server.SLoginRes_SelfHadLogin
		} else {
			err = errors.New("this gate id had connect")
			return
		}
	} else {
		if this.Id() != 0 {
			GSLog.Errorln("this.Id= ", this.Id())
			err = errors.New("this.Id != 0")
			return
		}

		this.SetId(vTClientId)
		err = this.UserMgr().addClient(this)
		if err != nil {
			return
		}
	}

	this.send_CmdServer2Gate_GetClientMsgRangeRes()
	this.send_CmdServer2Gate_LoginRes(vResult)
	return
}
