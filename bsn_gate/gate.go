package bsn_gate

import ()

type TGateId uint32

// IGate
type SGate struct {
	M_SServerUserMgr *SServerUserMgr
	M_SClientUserMgr *SClientUserMgr
	M_TGateId        TGateId
}

func newGate(vTGateId TGateId) (*SGate, error) {
	GLog.Debugln("newGate() vTGateId=", vTGateId)
	this := &SGate{
		M_TGateId: vTGateId,
	}

	var err error
	this.M_SServerUserMgr, err = newServerUserMgr()
	if err != nil {
		GLog.Errorln("newServerUserMgr fail")
		return nil, err
	}

	this.M_SClientUserMgr, err = newClientUserMgr()
	if err != nil {
		GLog.Errorln("newClientUserMgr fail")
		return nil, err
	}

	return this, nil
}

func (this *SGate) GetServerMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SGate) GetClientMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SGate) Listen() {
	this.GetServerMgr().Listen()
	this.GetClientMgr().Listen()
}

func (this *SGate) StopListen() {
	this.GetServerMgr().StopListen()
	this.GetClientMgr().StopListen()
}

func (this *SGate) Close() {
	this.GetServerMgr().Close()
	this.GetClientMgr().Close()
}
