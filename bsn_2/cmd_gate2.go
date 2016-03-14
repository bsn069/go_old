package bsn_2

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	bsn_gate "github.com/bsn069/go/bsn_gate2"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
)

type TId2Gate2 map[bsn_common.TGateId]*bsn_gate.SGate
type SCmdGate2 struct {
	M_TId2Gate TId2Gate2
	M_TGateId  bsn_common.TGateId
}

func NewCmdGate2() *SCmdGate2 {
	this := &SCmdGate2{
		M_TId2Gate: make(TId2Gate2),
		M_TGateId:  0,
	}
	return this
}

func (this *SCmdGate2) GATE2_RUN(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 3 {
		GSLog.Errorln("gateid clientListenPort")
		return
	}

	vuGateId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vuClientListenPort, err := strconv.ParseUint(vTInputParams[1], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSGate, err := this.Gate2Create(bsn_common.TGateId(vuGateId))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = vSGate.GetClientMgr().SetListenPort(uint16(vuClientListenPort))
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSGate.Run()
}

func (this *SCmdGate2) Gate2Create(vTGateId bsn_common.TGateId) (*bsn_gate.SGate, error) {
	vSGate, err := bsn_gate.NewGate(vTGateId)
	if err != nil {
		return nil, err
	}

	this.M_TId2Gate[vTGateId] = vSGate
	return vSGate, nil
}
