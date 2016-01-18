package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_net"
	"net"
)

type sClientMgr struct {
	m_Listener bsn_net.IListen
	m_clientId uint32
}

func newClientMgr() (*sClientMgr, error) {
	this := &sClientMgr{}
	this.m_Listener = bsn_net.NewListen()
	if this.m_Listener == nil {
		return nil, errors.New("create listener fail")
	}
	err := this.m_Listener.SetListenFunc(this.onListen())
	if err != nil {
		return nil, err
	}
	return this, nil
}

func (this *sClientMgr) SetListenPort(u16Port uint16) error {
	return this.m_Listener.SetListenPort(u16Port)
}

func (this *sClientMgr) Listen() error {
	return this.m_Listener.Listen()
}

func (this *sClientMgr) StopListen() {
	this.m_Listener.StopListen()
}

func (this *sClientMgr) onListen() bsn_net.TFuncOnListen {
	return func(conn net.Conn) {
		this.m_clientId++
		GLog.Debugln("this.m_clientId=", this.m_clientId)
		conn.Close()
	}
}
