package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_net"
	"net"
)

type sUser struct {
	m_iUserMgr IUserMgr
	m_userId   TUserId
	m_iConn    net.Conn
}

func newUser(userId TUserId, iConn net.Conn) (*sUser, error) {
	this := &sClient{m_userId: userId, m_iConn: iConn}
	return this, nil
}

func (this *sUser) GetId() TUserId {
	return this.m_u32Id
}

func (this *sUser) GetUserMgr() IUserMgr {
	return this.m_iUserMgr
}
