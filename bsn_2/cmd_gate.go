package bsn_2

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_gate"
	"net"
	"strconv"
	"sync"
	// "time"
	// "os"
)

type SCmdGate struct {
	M_TGateId2Gate    bsn_common.TGateId2Gate
	M_MutexGateCreate sync.Mutex
	M_TGateId         bsn_common.TGateId
	M_Listener        net.Listener
}

func NewCmdGate() *SCmdGate {
	this := &SCmdGate{
		M_TGateId2Gate: make(bsn_common.TGateId2Gate),
		M_TGateId:      0,
	}
	return this
}

func (this *SCmdGate) Test() {
	this.M_TGateId++

	vSGate, err := this.CreateGate(this.M_TGateId)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vstrAddr := ":" + strconv.Itoa(40000+int(this.M_TGateId))
	err = vSGate.GetClientMgr().SetListenAddr(vstrAddr)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = vSGate.GetClientMgr().Listen()
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	go func() {
		bsn_common.SleepSec(2)
		err := vSGate.GetClientMgr().StopListen()
		if err != nil {
			GSLog.Errorln(err)
			return
		}
	}()
}

func (this *SCmdGate) GetGate(vTGateId bsn_common.TGateId) (*bsn_gate.SGate, error) {
	if vTVoid, vbOk := this.M_TGateId2Gate[vTGateId]; vbOk {
		if vSGate, vbOk := vTVoid.(*bsn_gate.SGate); vbOk {
			return vSGate, nil
		} else {
			return nil, errors.New("type not is gate")
		}
	} else {
		return nil, errors.New("not found gate by id " + strconv.Itoa(int(vTGateId)))
	}
}

func (this *SCmdGate) CreateGate(vTGateId bsn_common.TGateId) (*bsn_gate.SGate, error) {
	this.M_MutexGateCreate.Lock()
	defer this.M_MutexGateCreate.Unlock()

	vSGate, err := this.GetGate(vTGateId)
	if err == nil {
		return nil, errors.New("have exist gate id")
	}

	vSGate, err = bsn_gate.NewGate(vTGateId)
	if err != nil {
		return nil, err
	}

	this.M_TGateId2Gate[vTGateId] = vSGate
	return vSGate, nil
}
