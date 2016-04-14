package bsn_client1

import (
// "errors"
// "github.com/bsn069/go/bsn_common"
// "github.com/bsn069/go/bsn_msg"
// "github.com/bsn069/go/bsn_net"
// "unsafe"
// "net"
// "sync"
)

func (this *SServerUserGate) ProcMsg_Gate_Echo() error {
	GSLog.Debugln("ProcMsg_Gate_Echo")
	GSLog.Debugln(string(this.M_by2MsgBody))
	return nil
}

func (this *SServerUserGate) ProcMsg_Gate_Pong() error {
	GSLog.Debugln("ProcMsg_Gate_Pong")
	return nil
}
