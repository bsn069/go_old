package bsn_2

// import (
// 	"errors"
// 	"github.com/bsn069/go/bsn_common"
// 	"github.com/bsn069/go/bsn_gate"
// 	"net"
// 	"strconv"
// 	"sync"
// 	// "time"
// 	// "os"
// )

// type SCmdGate struct {
// 	M_TGateId2Gate    bsn_common.TGateId2Gate
// 	M_MutexGateCreate sync.Mutex
// 	M_TGateId         bsn_common.TGateId
// 	M_Listener        net.Listener
// }

// func NewCmdGate() *SCmdGate {
// 	this := &SCmdGate{
// 		M_TGateId2Gate: make(bsn_common.TGateId2Gate),
// 		M_TGateId:      0,
// 	}
// 	return this
// }

// func (this *SCmdGate) GATE_CREATE(vTInputParams bsn_common.TInputParams) {
// 	if len(vTInputParams) != 1 {
// 		GSLog.Errorln("param must a number with gate id")
// 		return
// 	}

// 	vuGateId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	_, err = this.GateCreate(bsn_common.TGateId(vuGateId))
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}
// }

// func (this *SCmdGate) GATE_RUN(vTInputParams bsn_common.TInputParams) {
// 	if len(vTInputParams) != 3 {
// 		GSLog.Errorln("gateid clientListenPort serverListenPort")
// 		return
// 	}

// 	vuGateId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vuClientListenPort, err := strconv.ParseUint(vTInputParams[1], 10, 32)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vuServerListenPort, err := strconv.ParseUint(vTInputParams[2], 10, 32)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vSGate, err := this.GateCreate(bsn_common.TGateId(vuGateId))
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	err = vSGate.GetServerMgr().SetListenPort(uint16(vuServerListenPort))
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	err = vSGate.GetClientMgr().SetListenPort(uint16(vuClientListenPort))
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vSGate.Run()
// }

// func (this *SCmdGate) GATE_TEST(vTInputParams bsn_common.TInputParams) {
// 	this.M_TGateId++

// 	vSGate, err := this.GateCreate(this.M_TGateId)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vstrAddr := ":" + strconv.Itoa(40000+int(this.M_TGateId))
// 	err = vSGate.GetClientMgr().SetListenAddr(vstrAddr)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	err = vSGate.GetClientMgr().Listen()
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	go func() {
// 		bsn_common.SleepSec(2)
// 		err := vSGate.GetClientMgr().StopListen()
// 		if err != nil {
// 			GSLog.Errorln(err)
// 			return
// 		}
// 	}()
// }

// func (this *SCmdGate) GateGet(vTGateId bsn_common.TGateId) (*bsn_gate.SGate, error) {
// 	if vTVoid, vbOk := this.M_TGateId2Gate[vTGateId]; vbOk {
// 		if vSGate, vbOk := vTVoid.(*bsn_gate.SGate); vbOk {
// 			return vSGate, nil
// 		} else {
// 			return nil, errors.New("type not is gate")
// 		}
// 	} else {
// 		return nil, errors.New("not found gate by id " + strconv.Itoa(int(vTGateId)))
// 	}
// }

// func (this *SCmdGate) GateCreate(vTGateId bsn_common.TGateId) (*bsn_gate.SGate, error) {
// 	this.M_MutexGateCreate.Lock()
// 	defer this.M_MutexGateCreate.Unlock()

// 	vSGate, err := this.GateGet(vTGateId)
// 	if err == nil {
// 		return nil, errors.New("have exist gate id")
// 	}

// 	vSGate, err = bsn_gate.NewGate(vTGateId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	this.M_TGateId2Gate[vTGateId] = vSGate
// 	return vSGate, nil
// }
