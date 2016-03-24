package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	// "net"
	"sync/atomic"
)

type SNetServer struct {
	SNetListener

	M_chanNotifyClose chan bool
	M_i32Run          int32
}

func NewSNetServer() (*SNetServer, error) {
	GSLog.Debugln("NewSNetServer")
	this := &SNetServer{
		M_chanNotifyClose: make(chan bool, 1),
		M_i32Run:          0,
	}

	return this, nil
}

func (this *SNetServer) ShowInfo() {
	GSLog.Mustln("is running : ", this.M_i32Run == 1)
}

func (this *SNetServer) Run() error {
	if !atomic.CompareAndSwapInt32(&this.M_i32Run, 0, 1) {
		return errors.New("had listen")
	}

	this.Listen()
	go this.runImp()
	return nil
}

func (this *SNetServer) Close() error {
	if !atomic.CompareAndSwapInt32(&this.M_i32Run, 1, 2) {
		return errors.New("not listen")
	}
	GSLog.Debugln("close")

	// clear chan
	select {
	case <-this.M_chanNotifyClose:
	default:
	}
	this.M_chanNotifyClose <- true

	return nil
}

func (this *SNetServer) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		if atomic.CompareAndSwapInt32(&this.M_i32Run, 2, 0) {
			GSLog.Debugln("close complete")
		}
	}()

	GSLog.Debugln("run imp")
	var vbQuit bool = false
	for !vbQuit {
		vbQuit = true
		select {
		case vConn, ok := <-this.M_chanConn:
			if !ok {
				if atomic.CompareAndSwapInt32(&this.M_i32Run, 1, 2) {
					GSLog.Debugln("close from listen fail")
				}
				GSLog.Errorln("!ok")
				break
			}
			err := this.newConn()
			if err != nil {
				if atomic.CompareAndSwapInt32(&this.M_i32Run, 1, 2) {
					GSLog.Debugln("close from new user fail")
				}
				GSLog.Errorln(err)
				vConn.Close()
				break
			}

			vbQuit = false
		case <-this.M_chanNotifyClose:
			GSLog.Mustln("receive a notify to close")
		}
	}
}

func (this *SNetServer) newConn() error {
	return nil
}
