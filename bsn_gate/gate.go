package bsn_gate

import ()

type sGate struct {
	m_serverMgr IServerMgr
}

func newGate() IGate {
	var err error
	this := &sGate{}

	this.m_serverMgr, err = newServerMgr()
	if err != nil {
		GLog.Errorln("newServerMgr fail")
		return nil
	}

	return this
}

func (this *sGate) GetServerMgr() IServerMgr {
	return this.m_serverMgr
}
