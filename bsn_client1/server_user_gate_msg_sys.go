package bsn_client1

import (
	"bsn_define"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"errors"
	// "github.com/golang/protobuf/proto"
	"fmt"
)

func (this *SServerUserGate) procSysMsg(msgType bsn_define.ECmd) error {
	GSLog.Debugln("msgType=", msgType)

	switch msgType {
	case bsn_define.ECmd_Cmd_Ping:
		return this.ProcMsg_Cmd_Ping()
	case bsn_define.ECmd_Cmd_Pong:
		return this.ProcMsg_Cmd_Pong()
	}

	return errors.New(fmt.Sprintf("unknown msg type = %v", msgType))
}

func (this *SServerUserGate) ProcMsg_Cmd_Ping() (err error) {
	GSLog.Debugln("ProcMsg_Cmd_Ping")

	return this.Pong(this.M_by2MsgBody)
}

func (this *SServerUserGate) ProcMsg_Cmd_Pong() (err error) {
	GSLog.Debugln("ProcMsg_Cmd_Pong")
	return nil
}
