package bsn_net

import (
	"net"
)

type sUser struct {
	m_iUserMgr IUserMgr
	m_userId   TUserId
	m_iConn    net.Conn
}

func newUser(userId TUserId, iConn net.Conn) (*sUser, error) {
	this := &sUser{m_userId: userId, m_iConn: iConn}
	return this, nil
}

func (this *sUser) GetId() TUserId {
	return this.m_userId
}

func (this *sUser) GetUserMgr() IUserMgr {
	return this.m_iUserMgr
}

func (this *sUser) Close() {
	this.m_iConn.Close()
	this.m_iUserMgr.DelUser(this.GetId())
}
