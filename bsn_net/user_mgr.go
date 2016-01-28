package bsn_net

import (
	"errors"
	"net"
	// "sync"
)

type SUserMgr struct {
	IListen
	m_userId           TUserId
	m_id2User          TID2User
	m_iUserMgrCallBack IUserMgrCallBack
	m_type             TUserMgrType
}

func newUserMgr(userMgrType TUserMgrType, iUserMgrCallBack IUserMgrCallBack) (IUserMgr, error) {
	this := &sUserMgr{
		m_id2User:          make(TID2User, 0),
		m_iUserMgrCallBack: iUserMgrCallBack,
		m_type:             userMgrType,
	}
	this.IListen = NewListen(this)
	if this.IListen == nil {
		return nil, errors.New("create listener fail")
	}

	return this, nil
}

func (this *sUserMgr) GetType() TUserMgrType {
	return this.m_type
}

func (this *sUserMgr) GetUser(userId TUserId) IUser {
	if userImp, ok := this.m_id2User[userId]; ok {
		return userImp
	}
	return nil
}

func (this *sUserMgr) DelUser(userId TUserId) {
	GLog.Debugln("DelUser", userId)
	delete(this.m_id2User, userId)
}

func (this *sUserMgr) Close() {
	this.StopListen()

	for _, iUser := range this.m_id2User {
		iUser.Close()
	}
}

func (this *sUserMgr) OnListen(iConn net.Conn) error {
	this.m_userId++
	GLog.Debugln("OnListen this.m_userId=", this.m_userId)

	iUser, err := this.m_iUserMgrCallBack.NewUser(this.m_userId, iConn)
	if err != nil {
		GLog.Errorln(err)
		return err
	}

	this.m_id2User[iUser.GetId()] = iUser
	return nil
}
