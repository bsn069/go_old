package bsn_gate2

import (
	// "errors"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "time""
	// "net"
)

type SServerUserMgr struct {
}

func NewSServerUserMgr(vSGate *SGate) (*SServerUserMgr, error) {
	GSLog.Debugln("NewSServerUserMgr")
	this := &SServerUserMgr{}

	return this, nil
}

func (this *SServerUserMgr) Send(vSClientUser *SClientUser, vSMsgHeader *bsn_msg.SMsgHeader, vbyMsgBody []byte) error {
	return nil
}
