package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
	"net"
)

type sServerMgr struct {
	bsn_net.IUserMgr
}

func newServerMgr() (IServerMgr, error) {
	this := &sServerMgr{}
	var err error
	this.IUserMgr, err = bsn_net.NewUserMgr(this)
	if err != nil {
		return nil, err
	}
	return this, nil
}

func (this *sServerMgr) NewUser(userId bsn_net.TUserId, iConn net.Conn) (iUser bsn_net.IUser, err error) {
	iUser, err = bsn_net.NewUser(userId, iConn)
	return
}
