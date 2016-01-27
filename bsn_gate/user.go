package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
)

type sUser struct {
	bsn_net.IUser
}

func newUser(iUser bsn_net.IUser) (*sUser, error) {
	this := &sUser{}
	this.IUser = iUser
	return this, nil
}
