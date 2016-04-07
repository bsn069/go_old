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
	*SNetListener
	*bsn_common.SState
	*bsn_common.SNotifyClose

	M_INetServerImp INetServerImp
}

func NewSNetServer(vINetServerImp INetServerImp) (*SNetServer, error) {
	GSLog.Debugln("NewSNetServer")
	this := &SNetServer{}
	this.M_INetServerImp = vINetServerImp
	this.SNetListener, _ = NewSNetListener()
	this.SState = bsn_common.NewSState()
	this.SNotifyClose = bsn_common.NewSNotifyClose()
	return this, nil
}

func (this *SNetServer) Uninit() {
	this.M_INetServerImp = nil
}

func (this *SNetServer) ShowInfo() {
	GSLog.Mustln("ServerState : ", this.SState.M_TState.String())
}

func (this *SNetServer) Run() (err error) {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	this.SNotifyClose.Clear()

	this.Listen()
	go this.runImp()
	return nil
}

func (this *SNetServer) Close() (err error) {
	if !this.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		return errors.New("not listen")
	}
	GSLog.Debugln("close")

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	}()

	this.SNotifyClose.NotifyClose()
	this.SNotifyClose.WaitClose()

	return nil
}

func (this *SNetServer) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.StopListen()

		this.M_INetServerImp.NetServerImpOnClose()
		this.SNotifyClose.Close()
		this.Set(bsn_common.CState_Idle)
	}()

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	GSLog.Debugln("run imp")
	vbQuit := false
	for !vbQuit {
		vbQuit = true
		select {
		case vConn, ok := <-this.M_chanConn:
			if !ok {
				GSLog.Errorln("!ok")
				break
			}
			err := this.M_INetServerImp.NetServerImpAccept(vConn)
			if err != nil {
				GSLog.Errorln(err)
				vConn.Close()
				break
			}

			vbQuit = false
		case <-this.SNotifyClose.M_chanNotifyClose:
			GSLog.Mustln("receive a notify to close")
		}
	}
}
