package bsn_gate

import (
	"github.com/bsn069/go/bsn_net"
	// "net"
)

type TUserId2User map[TUserId]IUser
type TUserMgrType uint32

type IUserMgr interface {
	AddUser(vIUser IUser)
	DelUser(vIUser IUser)
	GetUser(vTUserId TUserId) IUser
	GetType() TUserMgrType
}

type SUserMgr struct {
	M_TUserMgrType TUserMgrType
	M_TUserId      TUserId
	M_users        TUserId2User
	*bsn_net.SListen
	M_bClose bool
}

func newUserMgr() (this *SUserMgr, err error) {
	GLog.Debugln("newUserMgr")
	this = &SUserMgr{
		M_users:   make(TUserId2User),
		M_TUserId: 0,
	}
	this.SListen = bsn_net.NewListen()
	return this, nil
}

func (this *SUserMgr) GetType() TUserMgrType {
	return this.M_TUserMgrType
}

func (this *SUserMgr) GenId() (TUserId, error) {
	this.M_TUserId++
	return this.M_TUserId, nil
}

func (this *SUserMgr) AddUser(vIUser IUser) {
	this.M_users[vIUser.GetId()] = vIUser
}

func (this *SUserMgr) DelUser(vIUser IUser) {
	delete(this.M_users, vIUser.GetId())
}

func (this *SUserMgr) GetUser(vTUserId TUserId) IUser {
	return this.M_users[vTUserId]
}
