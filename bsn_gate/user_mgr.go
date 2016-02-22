package bsn_gate

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_net"
	// "net"
)

type SUserMgr struct {
	M_TUserMgrType bsn_common.TGateUserMgrType
	M_TUserId      bsn_common.TGateUserId
	M_users        bsn_common.TGateUserId2User
	*bsn_net.SListen
	M_bClose bool
}

func newUserMgr() (this *SUserMgr, err error) {
	GSLog.Debugln("newUserMgr")
	this = &SUserMgr{
		M_users:   make(bsn_common.TGateUserId2User),
		M_TUserId: 0,
	}
	this.SListen = bsn_net.NewListen()
	return this, nil
}

func (this *SUserMgr) GetType() bsn_common.TGateUserMgrType {
	return this.M_TUserMgrType
}

func (this *SUserMgr) GenId() (bsn_common.TGateUserId, error) {
	this.M_TUserId++
	return this.M_TUserId, nil
}

func (this *SUserMgr) AddUser(vIUser bsn_common.IGateUser) {
	this.M_users[vIUser.GetId()] = vIUser
}

func (this *SUserMgr) DelUser(vIUser bsn_common.IGateUser) {
	delete(this.M_users, vIUser.GetId())
}

func (this *SUserMgr) GetUser(vTUserId bsn_common.TGateUserId) bsn_common.IGateUser {
	return this.M_users[vTUserId]
}
