package bsn_gate

import ()

type sGate struct {
	m_clientMgr IClientMgr
}

func newGate() IGate {
	this := &sGate{}
	var err error
	this.m_clientMgr, err = newClientMgr()
	if err != nil {
		GLog.Errorln("newClientMgr fail")
		return nil
	}
	return this
}

func (this *sGate) GetClientMgr() IClientMgr {
	return this.m_clientMgr
}
