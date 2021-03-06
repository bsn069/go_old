package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	"net"
	// "sync"
)

type SNetListener struct {
	*bsn_common.SState
	*bsn_common.SNotifyClose
	M_Listener net.Listener
	M_strAddr  string
	M_chanConn bsn_common.TNetChanConn
}

func NewSNetListener() (*SNetListener, error) {
	GSLog.Debugln("NewSNetListener")
	this := &SNetListener{
		M_chanConn: make(bsn_common.TNetChanConn, 100),
	}
	this.SState = bsn_common.NewSState()
	this.SNotifyClose = bsn_common.NewSNotifyClose()

	return this, nil
}

func (this *SNetListener) ShowInfo() {
	GSLog.Mustln("listen addr: ", this.M_strAddr)
	GSLog.Mustln("is listen  : ", this.IsListen())
}

func (this *SNetListener) SetAddr(strAddr string) error {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}
	defer this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)

	if strAddr == "" {
		return errors.New("error addres")
	}
	this.M_strAddr = strAddr
	return nil
}

func (this *SNetListener) Addr() string {
	return this.M_strAddr
}

func (this *SNetListener) IsListen() bool {
	return this.Is(bsn_common.CState_Runing)
}

func (this *SNetListener) Listen() (err error) {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	if this.Addr() == "" {
		err = errors.New("no address")
		return
	}

	GSLog.Mustln("listen strListenAddr ", this.Addr())
	this.M_Listener, err = net.Listen("tcp", this.Addr())
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	this.SNotifyClose.Clear()

	go this.listenFunc()
	return
}

func (this *SNetListener) StopListen() (err error) {
	if !this.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		return errors.New("not listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	}()

	// GSLog.Debugln("3")
	err = this.M_Listener.Close()
	if err != nil {
		GSLog.Debugln("4" + err.Error())
	}
	this.M_Listener = nil

	this.SNotifyClose.NotifyClose()
	this.SNotifyClose.WaitClose()

	return nil
}

func (this *SNetListener) listenFunc() {
	GSLog.Mustln("listenFunc")
	vListener := this.M_Listener
	defer bsn_common.FuncGuard()
	defer func() {
		// GSLog.Debugln("send close before")

		GSLog.Debugln("close all connect")
		for vConn := range this.M_chanConn {
			vConn.Close()
		}
		// GSLog.Debugln("send close after")

		this.SNotifyClose.Close()
		this.Set(bsn_common.CState_Idle)
	}()

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)

	vbQuit := false
	for !vbQuit {
		vConn, err := vListener.Accept()
		if err != nil {
			GSLog.Errorln(err)
			vbQuit = true
			continue
		}

		select {
		case this.M_chanConn <- vConn:
		case <-this.SNotifyClose.M_chanNotifyClose:
			vbQuit = true
		}
	}
}
