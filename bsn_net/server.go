package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	"net"
	// "sync/atomic"
)

type INetServerImp interface {
	NetServerImpAccept(vConn net.Conn) error
	NetServerImpOnClose() error
}

type SNetServer struct {
	SNetListener

	M_chanNotifyClose chan bool
	M_SState          *bsn_common.SState

	M_INetServerImp INetServerImp
}

func NewSNetServer() (*SNetServer, error) {
	GSLog.Debugln("NewSNetServer")
	this := &SNetServer{
		M_chanNotifyClose: make(chan bool, 1),
		M_SState:          bsn_common.NewSState(),
	}

	return this, nil
}

func (this *SNetServer) Uninit() {
	this.M_INetServerImp = nil
}

func (this *SNetServer) ShowInfo() {
	GSLog.Mustln("ServerState : ", this.M_SState.M_TState.String())
}

func (this *SNetServer) Run() error {
	if !this.M_SState.Change(bsn_common.CState_Idle, bsn_common.CState_Runing) {
		return errors.New("had listen")
	}

	this.Listen()
	go this.runImp()
	return nil
}

func (this *SNetServer) Close() error {
	if !this.M_SState.Change(bsn_common.CState_Runing, bsn_common.CState_Stoping) {
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
		this.StopListen()

		if this.M_SState.Change(bsn_common.CState_Stoping, bsn_common.CState_Idle) {
			GSLog.Debugln("close complete")
			this.M_INetServerImp.NetServerImpOnClose()
		}
	}()

	GSLog.Debugln("run imp")
	vbQuit := false
	for !vbQuit {
		vbQuit = true
		select {
		case vConn, ok := <-this.M_chanConn:
			if !ok {
				if this.M_SState.Change(bsn_common.CState_Runing, bsn_common.CState_Stoping) {
					GSLog.Debugln("close from listen fail")
				}
				GSLog.Errorln("!ok")
				break
			}
			err := this.M_INetServerImp.NetServerImpAccept(vConn)
			if err != nil {
				if this.M_SState.Change(bsn_common.CState_Runing, bsn_common.CState_Stoping) {
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
