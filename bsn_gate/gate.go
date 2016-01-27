package bsn_gate

import ()

type sGate struct {
	m_serverMgr IUserMgr
	m_clientMgr IUserMgr
}

func newGate() IGate {
	var err error
	this := &sGate{}

	this.m_serverMgr, err = newUserMgr(CClientMgr)
	if err != nil {
		GLog.Errorln("newUserMgr fail")
		return nil
	}

	this.m_clientMgr, err = newUserMgr(CServerMgr)
	if err != nil {
		GLog.Errorln("newUserMgr fail")
		return nil
	}

	return this
}

func (this *sGate) GetServerMgr() IUserMgr {
	return this.m_serverMgr
}

func (this *sGate) GetClientMgr() IUserMgr {
	return this.m_clientMgr
}

func (this *sGate) Listen() {
	this.GetClientMgr().Listen()
	this.GetServerMgr().Listen()
}

func (this *sGate) StopListen() {
	this.GetClientMgr().StopListen()
	this.GetServerMgr().StopListen()
}

func (this *sGate) Close() {
	this.GetClientMgr().Close()
	this.GetServerMgr().Close()
}
