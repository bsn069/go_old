package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
	"net"
)

// bsn_net.IUserMgr
// bsn_net.IUserMgrCallBack
type sUserMgr struct {
	bsn_net.IUserMgr
}

func newUserMgr(vTUserMgrType bsn_net.TUserMgrType) (this *sUserMgr, err error) {
	GLog.Debugln("newUserMgr", vTUserMgrType)
	this = &sUserMgr{}
	this.IUserMgr, err = bsn_net.NewUserMgr(vTUserMgrType, this)
	if err != nil {
		return nil, err
	}
	return this, nil
}

// bsn_net.IUserMgrCallBack
func (this *sUserMgr) NewUser(vTUserId bsn_net.TUserId, vConn net.Conn) (bsn_net.IUser, error) {
	vnetIUser, err := bsn_net.NewUser(this, vTUserId, vConn)
	if err != nil {
		return nil, err
	}

	var vsUser *sUser
	if this.GetType() == CClientMgr {
		vsClientUser, err := newClientUser(vnetIUser)
		vsUser, _ := vsClientUser.(*sUser)
	} else {
		vsUser, err = newClientUser(vnetIUser)
	}
	if err != nil {
		return nil, err
	}
	return vsUser, nil
}
