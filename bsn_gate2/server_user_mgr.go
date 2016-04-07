package bsn_gate2

import (
	// "errors"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "time""
	// "net"
)

type SServerUserMgr struct {
	M_SGate *SGate
	M_Users []*SServerUser
}

func NewSServerUserMgr(vSGate *SGate) (*SServerUserMgr, error) {
	GSLog.Debugln("NewSServerUserMgr")
	this := &SServerUserMgr{
		M_SGate: vSGate,
		M_Users: make([]*SServerUser, 1),
	}

	return this, nil
}

func (this *SServerUserMgr) ShowInfo() {
}

func (this *SServerUserMgr) Send(vSClientUser *SClientUser, vSMsgHeader *bsn_msg.SMsgHeader, vbyMsgBody []byte) error {
	GSLog.Debugln("Send")
	GSLog.Mustln(vSClientUser)
	GSLog.Mustln(vSMsgHeader)
	GSLog.Mustln(vbyMsgBody)
	GSLog.Mustln(string(vbyMsgBody))
	return nil
}

func (this *SServerUserMgr) Gate() *SGate {
	return this.M_SGate
}

func (this *SServerUserMgr) Run() error {
	// connect config server
	vUser, err := NewSServerUser(this, "localhost:50001")
	if err != nil {
		return err
	}

	err = vUser.Run()
	if err != nil {
		return err
	}
	this.M_Users = append(this.M_Users, vUser)

	return nil
}

func (this *SServerUserMgr) Close() error {
	return nil
}
