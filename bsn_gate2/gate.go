package bsn_gate2

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_log"
	"strconv"
)

type SGate struct {
	M_SClientUserMgr *SClientUserMgr
	M_TGateId        bsn_common.TGateId
}

func NewGate(vTGateId bsn_common.TGateId) (*SGate, error) {
	GSLog.Debugln("newGate() vTGateId=", vTGateId)
	this := &SGate{
		M_TGateId: vTGateId,
	}

	var err error

	this.M_SClientUserMgr, err = NewSClientUserMgr(this)
	if err != nil {
		GSLog.Errorln("newClientUserMgr fail")
		return nil, err
	}

	vSCmd := &SCmd{M_SGate: this}
	bsn_input.GSInput.Reg("Gate2_"+strconv.Itoa(int(vTGateId)), vSCmd)

	return this, nil
}

func (this *SGate) ShowInfo() {
	GSLog.Mustln("ClientMgr")
	this.GetClientMgr().ShowInfo()
}

func (this *SGate) GetClientMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SGate) Close() {
	this.GetClientMgr().Close()
}

func (this *SGate) Run() {
	this.GetClientMgr().Run()
}
