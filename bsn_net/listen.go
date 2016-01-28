package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"net"
	"strconv"
	"sync"
)

type IListenVirtual interface {
	OnListen(conn net.Conn) error
}

type SListen struct {
	IListenVirtual

	m_TPort TPort

	m_Listener   net.Listener
	m_Mutex      sync.Mutex
	m_TCloseChan TCloseChan
}

func newListen(vIListenVirtual IListenVirtual) *SListen {
	this := &sListen{
		IListenVirtual: vIListenVirtual,
		m_chanClose:    make(TCloseChan, 1),
	}
	return this
}

func (this *SListen) SetListenPort(vTPort TPort) error {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if vTPort == 0 {
		return errors.New("error tcp port")
	}
	if this.m_Listener != nil {
		return errors.New("had listen")
	}
	this.m_TPort = vTPort
	return nil
}

func (this *SListen) Listen() (err error) {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if this.m_TPort == 0 {
		return errors.New("error tcp port")
	}
	if this.m_Listener != nil {
		return errors.New("had listen")
	}

	strListenAddr := ":" + strconv.Itoa((int)(this.m_TPort))
	GLog.Mustln("listen strListenAddr", strListenAddr)

	this.m_Listener, err = net.Listen("tcp", strListenAddr)
	if err != nil {
		return
	}

	// if not call StopListen, clear close chanel
	select {
	case <-this.m_TCloseChan:
	default:
	}

	go this.listenFunc()
	return
}

// block until close
func (this *SListen) StopListen() {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if this.m_Listener == nil {
		return
	}
	this.m_Listener.Close()
	<-this.m_TCloseChan
}

func (this *SListen) listenFunc() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.m_Listener.Close()
		this.m_Listener = nil
		this.m_TCloseChan <- true
	}()

	for {
		vConn, err := this.m_Listener.Accept()
		if err != nil {
			GLog.Errorln(err)
			return
		}

		err = this.OnListen(vConn)
		if err != nil {
			GLog.Errorln(err)
			vConn.Close()
		}
	}
}
