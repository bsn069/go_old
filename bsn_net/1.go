package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	// "fmt"
	"net"
	// "sync/atomic"
	"sync"
)

type IRecverCB interface {
	RecverCBOnAcceptNetConn(vConn net.Conn) error
	LRecverCBOnClose() error
}

type STCPRecver struct {
	M_RWMutex      sync.RWMutex
	M_IRecverCB    IRecverCB
	M_SNotifyClose bsn_common.SNotifyClose
	M_SState       bsn_common.SState

	M_Conn net.Conn
	M_Id   uint16
}

func NewSTCPRecver() (this *STCPRecver, err error) {
	GSLog.Debugln("NewSTCPRecver")
	this = &STCPRecver{}
	this.Reset()

	return this, nil
}

func (this *STCPRecver) Reset() (err error) {
	this.M_RWMutex.Lock()
	defer this.M_RWMutex.Unlock()

	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("not idle")
	}

	this.M_SState.Reset()
	this.M_SNotifyClose.Reset()
	this.M_IRecverCB = nil
	this.M_Id = 0
	this.M_Conn = nil

	return nil
}

func (this *STCPRecver) SetIListenerCB(vIRecverCB IRecverCB) (err error) {
	this.M_RWMutex.Lock()
	defer this.M_RWMutex.Unlock()

	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}
	defer this.M_SState.Set(bsn_common.CState_Idle)

	if vIRecverCB == nil {
		return errors.New("vIRecverCB is nil")
	}

	this.M_IRecverCB = vIRecverCB
	return nil
}

func (this *STCPRecver) SetConn(vConn net.Conn) (err error) {
	this.M_RWMutex.Lock()
	defer this.M_RWMutex.Unlock()

	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had run")
	}
	defer this.M_SState.Set(bsn_common.CState_Idle)

	if vConn == nil {
		return errors.New("vConn == nil")
	}

	this.M_Conn = vConn
	return nil
}

func (this *STCPRecver) SetId(id uint16) (err error) {
	this.M_RWMutex.Lock()
	defer this.M_RWMutex.Unlock()

	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had run")
	}
	defer this.M_SState.Set(bsn_common.CState_Idle)

	if id == 0 {
		return errors.New("id == 0")
	}

	this.M_Id = id
	return nil
}

func (this *STCPRecver) Id() (id uint16) {
	this.M_RWMutex.RLock()
	defer this.M_RWMutex.RUnlock()

	return this.M_Id
}
