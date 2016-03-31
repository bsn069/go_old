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
	M_Listener        net.Listener
	M_strAddr         string
	M_chanConn        bsn_common.TNetChanConn
	M_chanNotifyClose chan bool
	M_SState          *bsn_common.SState
}

func NewSNetListener() (*SNetListener, error) {
	GSLog.Debugln("NewSNetListener")
	this := &SNetListener{
		M_chanConn:        make(bsn_common.TNetChanConn, 100),
		M_chanNotifyClose: make(chan bool, 1),
	}

	return this, nil
}

func (this *SNetListener) ShowInfo() {
	GSLog.Mustln("listen addr: ", this.M_strAddr)
	GSLog.Mustln("is listen  : ", this.IsListen())
}

func (this *SNetListener) SetAddr(strAddr string) error {
	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}
	defer this.M_SState.Change(bsn_common.CState_Op, bsn_common.CState_Idle)

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
	return this.M_SState.IsState(bsn_common.CState_Runing)
}

func (this *SNetListener) Listen() (err error) {
	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}
	defer this.M_SState.Change(bsn_common.CState_Op, bsn_common.CState_Idle)

	if this.Addr() == "" {
		return errors.New("no address")
	}

	GSLog.Mustln("listen strListenAddr ", this.Addr())
	this.M_Listener, err = net.Listen("tcp", this.Addr())
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	this.M_SState.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	go this.listenFunc()
	return
}

func (this *SNetListener) StopListen() error {
	if !this.IsListen() {
		return errors.New("not listen")
	}
	// GSLog.Debugln("3")
	err := this.M_Listener.Close()
	if err != nil {
		GSLog.Debugln("4" + err.Error())
	}
	this.M_Listener = nil

	select {
	case <-this.M_chanNotifyClose:
	default:
	}
	this.M_chanNotifyClose <- true

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
	}()

	vbQuit := false
	for !vbQuit {
		// GSLog.Debugln("wait accept")
		vConn, err := vListener.Accept()
		// GSLog.Debugln("have accept")
		if err != nil {
			GSLog.Errorln(err)
			vbQuit = true
			continue
		}

		// GSLog.Debugln("send conn before")
		select {
		case this.M_chanConn <- vConn:
		case <-this.M_chanNotifyClose:
			vbQuit = true
		}
		// GSLog.Debugln("send conn after")
	}
}
