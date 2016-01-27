package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
)

type sClientUser struct {
	*sUser
}

func newClientUser(iUser bsn_net.IUser) (IUser, error) {
	GLog.Debugln("newClientUser")
	vsUser, err := newUser(iUser)
	if err != nil {
		return nil, err
	}
	this := &sClientUser{}
	this.sUser = vsUser
	return this, nil
}
