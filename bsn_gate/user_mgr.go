package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
	"net"
)

type sUserMgr struct {
	bsn_net.IUserMgr
}

func newUserMgr(userMgrType bsn_net.TUserMgrType) (IUserMgr, error) {
	this := &sUserMgr{}
	var err error
	this.IUserMgr, err = bsn_net.NewUserMgr(userMgrType, this)
	if err != nil {
		return nil, err
	}
	return this, nil
}

func (this *sUserMgr) NewUser(userId bsn_net.TUserId, iConn net.Conn) (bsn_net.IUser, error) {
	netiUser, err := bsn_net.NewUser(this, userId, iConn)
	if err != nil {
		return nil, err
	}

	var iUser IUser
	if this.GetType() == CClientMgr {
		iUser, err = newClientUser(netiUser)
	} else {
		iUser, err = newServerUser(netiUser)
	}
	if err != nil {
		return nil, err
	}
	return iUser, nil
}
