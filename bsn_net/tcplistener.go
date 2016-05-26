package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_log"
	// "time"
	// "math"
	// "fmt"
	"net"
	"sync"
	// "sync/atomic"
)

type IListenerCB interface {
	ListenerCBOnAcceptNetConn(vConn net.Conn) error
	ListenerCBOnClose() error
}

type STCPListener struct {
	M_Listener      net.Listener
	M_strListenAddr string
	M_SNotifyClose  bsn_common.SNotifyClose
	M_SState        bsn_common.SState
	M_RWMutex       sync.RWMutex
	M_IListenerCB   IListenerCB
}

func NewSTCPListener() (this *STCPListener, err error) {
	GSLog.Debugln("NewSTCPListener")
	this = &STCPListener{}
	this.Reset()
	return
}

func (this *STCPListener) Reset() (err error) {
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
	this.M_strListenAddr = ":0"
	this.M_Listener = nil
	this.M_IListenerCB = nil

	return nil
}

func (this *STCPListener) SetIListenerCB(vIListenerCB IListenerCB) (err error) {
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

	if vIListenerCB == nil {
		return errors.New("vIListenerCB is nil")
	}

	this.M_IListenerCB = vIListenerCB
	return nil
}

func (this *STCPListener) SetAddr(strAddr string) (err error) {
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

	if strAddr == "" {
		return errors.New("strAddr is empty")
	}

	this.M_strListenAddr = strAddr
	return nil
}

func (this *STCPListener) Addr() (strAddr string) {
	this.M_RWMutex.RLock()
	defer this.M_RWMutex.RUnlock()

	return this.M_strListenAddr
}

func (this *STCPListener) Start() (err error) {
	this.M_RWMutex.Lock()
	defer this.M_RWMutex.Unlock()

	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("error state")
	}

	defer func() {
		if err != nil {
			this.M_SState.Set(bsn_common.CState_Idle)
		}
	}()

	if this.M_strListenAddr == "" {
		return errors.New("this.M_strListenAddr is empty")
	}

	if this.M_IListenerCB == nil {
		return errors.New("this.M_IListenerCB is nil")
	}

	GSLog.Mustln("listenTCP ", this.M_strListenAddr)
	this.M_Listener, err = net.Listen("tcp", this.M_strListenAddr)
	if err != nil {
		return err
	}

	this.M_SNotifyClose.Clear()
	go this.workerTCPListen()
	return nil
}

func (this *STCPListener) workerTCPListen() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.M_SNotifyClose.Close()
		this.M_Listener = nil
		this.M_SState.Set(bsn_common.CState_Idle)
		this.M_IListenerCB.ListenerCBOnClose()
	}()

	this.M_SState.Change(bsn_common.CState_Op, bsn_common.CState_Runing)

	vbQuit := false
	for !vbQuit {
		vConn, err := this.M_Listener.Accept()
		if err != nil {
			GSLog.Errorln(err)
			vbQuit = true
			continue
		}
		this.onTCPClient(vConn)
	}
}

func (this *STCPListener) onTCPClient(vConn net.Conn) (err error) {
	defer func() {
		if err != nil {
			GSLog.Errorln(err)
			vConn.Close()
		}
	}()

	vstrAddr := vConn.RemoteAddr().String()
	GSLog.Debugf("client connect addr=%v\n", vstrAddr)

	return this.M_IListenerCB.ListenerCBOnAcceptNetConn(vConn)
}

func (this *STCPListener) Stop() (err error) {
	this.M_RWMutex.Lock()
	defer this.M_RWMutex.Unlock()

	defer func() {
		if err != nil {
			GSLog.Errorln(err)
		}
	}()

	if !this.M_SState.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		return errors.New("error state")
	}

	defer func() {
		if err != nil {
			this.M_SState.Set(bsn_common.CState_Runing)
		}
	}()

	err = this.M_Listener.Close()
	if err != nil {
		return err
	}

	this.M_SNotifyClose.WaitClose()

	return nil
}
