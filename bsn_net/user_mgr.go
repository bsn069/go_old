package bsn_net

import (
	"errors"
	"net"
	"sync"
)

type sUserMgr struct {
	m_userId      TUserId
	m_id2User     TID2User
	m_iListener   IListen
	m_funcNewUser TFuncNewUser
	m_bClose      bool
	m_Mutex       sync.Mutex
}

func newUserMgr() (IUserMgr, error) {
	this := &sUserMgr{
		m_userId:  0,
		m_id2User: make(TID2User, 0),
		m_bClose:  false,
	}
	this.m_iListener = NewListen()
	if this.m_iListener == nil {
		return nil, errors.New("create listener fail")
	}
	err := this.m_iListener.SetListenFunc(this.onListen())
	if err != nil {
		return nil, err
	}

	this.m_funcNewUser = funcNewUser
	return this, nil
}

func funcNewUser(user IUser) (TUserImp, error) {
	return nil, errors.New("not new user func")
}

func (this *sUserMgr) SetFuncNewUser(funcNewUser TFuncNewUser) error {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if funcNewUser == nil {
		return errors.New("not func")
	}
	this.m_funcNewUser = funcNewUser
	return nil
}

func (this *sUserMgr) SetListenPort(port TPort) error {
	return this.m_iListener.SetListenPort(port)
}

func (this *sUserMgr) Listen() error {
	return this.m_iListener.Listen()
}

func (this *sUserMgr) StopListen() {
	this.m_iListener.StopListen()
}

func (this *sUserMgr) GetUser(userId TUserId) TUserImp {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if userImp, ok := this.m_id2User[userId]; ok {
		return userImp
	}
	return nil
}

func (this *sUserMgr) DelUser(userId TUserId) {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()
	s
	delete(this.m_id2User, userId)
}

func (this *sUserMgr) Close() {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	this.StopListen()
	s
	for _, userImp := range this.m_id2User {
		iUser, ok := userImp.(IUser)
		if ok {
			iUser.Close()
		}
	}
}

func (this *sUserMgr) onListen() TFuncOnListen {
	return func(iConn net.Conn) error {
		this.m_userId++
		GLog.Debugln("this.m_userId=", this.m_userId)

		iUser, err := newUser(this.m_userId, iConn)
		if err != nil {
			GLog.Errorln(err)
			return err
		}

		userImp, err := this.m_funcNewUser(iUser)
		if err != nil {
			GLog.Errorln(err)
			return err
		}

		this.m_id2User[iUser.GetId()] = userImp
		return nil
	}
}
