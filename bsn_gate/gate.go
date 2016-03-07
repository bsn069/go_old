package bsn_gate

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_log"
	"strconv"
)

type SGate struct {
	M_SServerUserMgr *SServerUserMgr
	M_SClientUserMgr *SClientUserMgr
	M_TGateId        bsn_common.TGateId
}

func NewGate(vTGateId bsn_common.TGateId) (*SGate, error) {
	GSLog.Debugln("newGate() vTGateId=", vTGateId)
	this := &SGate{
		M_TGateId: vTGateId,
	}

	var err error
	this.M_SServerUserMgr, err = NewServerUserMgr(this)
	if err != nil {
		GSLog.Errorln("newServerUserMgr fail")
		return nil, err
	}

	this.M_SClientUserMgr, err = NewClientUserMgr(this)
	if err != nil {
		GSLog.Errorln("newClientUserMgr fail")
		return nil, err
	}

	vSCmd := &SCmd{M_SGate: this}
	bsn_input.GSInput.Reg("Gate"+strconv.Itoa(int(vTGateId)), vSCmd)

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
