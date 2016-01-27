package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
)

type sServerUser struct {
	*sUser
}

func newServerUser(iUser bsn_net.IUser) (IUser, error) {
	GLog.Debugln("newServerUser")
	vsUser, err := newUser(iUser)
	if err != nil {
		return nil, err
	}
	this := &sServerUser{}
	this.sUser = vsUser
	return this, nil
}
