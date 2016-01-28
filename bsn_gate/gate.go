package bsn_gate

import ()

// IGate
type sGate struct {
	m_serverMgr *sUserMgr
	m_clientMgr *sUserMgr
}

func newGate() (IGate, err error) {
	GLog.Debugln("newGate()")
	this := &sGate{}

	this.m_serverMgr, err = newUserMgr(CClientMgr)
	if err != nil {
		GLog.Errorln("newUserMgr fail")
		return nil, err
	}

	this.m_clientMgr, err = newUserMgr(CServerMgr)
	if err != nil {
		GLog.Errorln("newUserMgr fail")
		return nil, err
	}

	return this, nil
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
