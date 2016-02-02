package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"net"
	// "strconv"
	"sync"
)

type SListen struct {
	m_strAddr string

	m_Listener  net.Listener
	m_Mutex     sync.Mutex
	m_chanClose TChanClose
	M_chanConn  TChanConn
}

func newListen() *SListen {
	this := &SListen{
		m_chanClose: make(TChanClose, 1),
		M_chanConn:  make(TChanConn, 100),
	}
	return this
}

func (this *SListen) SetListenAddr(strAddr string) error {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if strAddr == "" {
		return errors.New("error addres")
	}
	if this.m_Listener != nil {
		return errors.New("had listen")
	}
	this.m_strAddr = strAddr
	return nil
}

func (this *SListen) Listen() (err error) {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if this.m_strAddr == "" {
		return errors.New("no address")
	}
	if this.m_Listener != nil {
		return errors.New("had listen")
	}

	GLog.Mustln("listen strListenAddr ", this.m_strAddr)
	this.m_Listener, err = net.Listen("tcp", this.m_strAddr)
	if err != nil {
		GLog.Errorln(err)
		return
	}

	// if not call StopListen, clear close chanel
	select {
	case <-this.m_chanClose:
	default:
	}

	go this.listenFunc()
	return
}

// block until close
func (this *SListen) StopListen() error {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()
	GLog.Debugln("1")

	if this.m_Listener == nil {
		GLog.Debugln("2")
		return errors.New("not listen")
	}
	GLog.Debugln("3")
	err := this.m_Listener.Close()
	if err != nil {

		GLog.Debugln("4" + err.Error())
	}
	// close(this.M_chanConn)
	GLog.Debugln("5")
	<-this.m_chanClose
	GLog.Debugln("6")
	this.m_Listener = nil
	GLog.Debugln("7")
	return nil
}

func (this *SListen) listenFunc() {
	defer bsn_common.FuncGuard()
	defer func() {
		GLog.Debugln("send close before")
		this.m_chanClose <- true
		GLog.Debugln("send close before")
	}()

	GLog.Mustln("listenFunc")
	for {
		GLog.Debugln("wait accept")
		vConn, err := this.m_Listener.Accept()
		GLog.Debugln("have accept")
		if err != nil {
			GLog.Errorln(err)
			return
		}

		GLog.Debugln("send conn before")
		this.M_chanConn <- vConn
		GLog.Debugln("send conn after")
	}
}
