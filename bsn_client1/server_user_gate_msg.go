package bsn_client1

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	"errors"
	// "net"
	// "strconv"
	// "sync"
	"fmt"
)

func (this *SServerUserGate) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")

	switch this.MsgType() {
	case bsn_msg.GMsgDefine_Gate2Client_Ping:
		return this.ProcMsg_Gate_Ping()
	case bsn_msg.GMsgDefine_Gate2Client_Pong:
		return this.ProcMsg_Gate_Pong()
	case bsn_msg.GMsgDefine_Echo2Client_Ping:
		return this.ProcMsg_Echo_Ping()
	case bsn_msg.GMsgDefine_Echo2Client_Pong:
		return this.ProcMsg_Echo_Pong()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", this.M_SMsgHeader.Type())
	return errors.New(strInfo)
}
