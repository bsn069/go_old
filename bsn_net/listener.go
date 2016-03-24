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
	M_Listener net.Listener
	M_strAddr  string
	M_chanConn bsn_common.TNetChanConn
}

func NewSNetListener() (*SNetListener, error) {
	GSLog.Debugln("NewSNetListener")
	this := &SNetListener{
		M_chanConn: make(bsn_common.TNetChanConn, 100),
	}

	return this, nil
}

func (this *SNetListener) ShowInfo() {
	GSLog.Mustln("listen addr: ", this.M_strAddr)
	GSLog.Mustln("is listen  : ", this.IsListen())
}

func (this *SNetListener) SetAddr(strAddr string) error {
	if strAddr == "" {
		return errors.New("error addres")
	}
	if this.IsListen() {
		return errors.New("had listen")
	}
	this.M_strAddr = strAddr
	return nil
}

func (this *SNetListener) Addr() string {
	return this.M_strAddr
}

func (this *SNetListener) IsListen() bool {
	return this.M_Listener != nil
}

func (this *SNetListener) Listen() (err error) {
	// this.M_Mutex.Lock()
	// defer this.M_Mutex.Unlock()

	if this.Addr() == "" {
		return errors.New("no address")
	}
	if this.IsListen() {
		return errors.New("had listen")
	}

	GSLog.Mustln("listen strListenAddr ", this.Addr())
	this.M_Listener, err = net.Listen("tcp", this.Addr())
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	go this.listenFunc()
	return
}

func (this *SNetListener) StopListen() error {
	// this.M_Mutex.Lock()
	// defer this.M_Mutex.Unlock()
	// GSLog.Debugln("1")

	if !this.IsListen() {
		return errors.New("not listen")
	}
	// GSLog.Debugln("3")
	err := this.M_Listener.Close()
	if err != nil {
		GSLog.Debugln("4" + err.Error())
	}
	this.M_Listener = nil

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

	for {
		// GSLog.Debugln("wait accept")
		vConn, err := vListener.Accept()
		// GSLog.Debugln("have accept")
		if err != nil {
			GSLog.Errorln(err)
			return
		}

		// GSLog.Debugln("send conn before")
		this.M_chanConn <- vConn
		// GSLog.Debugln("send conn after")
	}
}
