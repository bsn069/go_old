package bsn_gate

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_net"
	// "net"
)

type SUserMgr struct {
	M_TUserMgrType bsn_common.TGateUserMgrType
	M_TUserId      bsn_common.TGateUserId
	M_TId2User     bsn_common.TGateUserId2User
	*bsn_net.SListen
	M_bClose bool
	M_SGate  *SGate
}

func NewUserMgr(vSGate *SGate) (this *SUserMgr, err error) {
	GSLog.Debugln("newUserMgr")
	this = &SUserMgr{
		M_TId2User: make(bsn_common.TGateUserId2User),
		M_TUserId:  0,
		M_SGate:    vSGate,
	}
	this.SListen = bsn_net.NewListen()
	return this, nil
}

func (this *SUserMgr) Gate() *SGate {
	return this.M_SGate
}

func (this *SUserMgr) Type() bsn_common.TGateUserMgrType {
	return this.M_TUserMgrType
}

func (this *SUserMgr) GenUserId() (bsn_common.TGateUserId, error) {
	this.M_TUserId++
	return this.M_TUserId, nil
}

func (this *SUserMgr) AddUser(vIUser bsn_common.IGateUser) error {
	this.M_TId2User[vIUser.Id()] = vIUser
	return nil
}

func (this *SUserMgr) DelUser(vIUser bsn_common.IGateUser) error {
	delete(this.M_TId2User, vIUser.Id())
	return nil
}

func (this *SUserMgr) User(vTUserId bsn_common.TGateUserId) (bsn_common.IGateUser, error) {
	return this.M_TId2User[vTUserId], nil
}
