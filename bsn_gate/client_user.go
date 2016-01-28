package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
)

type sClientUser struct {
	*sUser
}

func newClientUser(net_iUser bsn_net.IUser) (*sClientUser, error) {
	GLog.Debugln("newClientUser")
	this := &sClientUser{}

	iUser, err := newUser(net_iUser, this)
	if err != nil {
		return nil, err
	}
	this.IUser = iUser
	return this, nil
}

func (this *sClientUser) GetMsgBodyLen() uint {
	GLog.Debugln("GetMsgBodyLen")
	return 6
}

func (this *sClientUser) OnReadOneMsg() error {
	GLog.Debugln("OnReadOneMsg")
	GLog.Debugln(this.m_byMsgHeader)
	GLog.Debugln(this.m_byMsgBody)
	return nil
}

func (this *sClientUser) OnDisconnect() {
	GLog.Debugln("OnDisconnect")
}
