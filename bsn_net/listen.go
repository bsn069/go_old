package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"net"
	"strconv"
	"sync"
)

type sListen struct {
	m_u16Port uint16

	m_Listener        net.Listener
	m_MutexStopListen sync.Mutex
	m_funcOnListen    TFuncOnListen
}

func newListen() *sListen {
	this := &sListen{}
	return this
}

func (this *sListen) SetListenFunc(funcOnListen TFuncOnListen) error {
	if this.m_Listener != nil {
		return errors.New("had listen")
	}
	this.m_funcOnListen = funcOnListen
	return nil
}

func (this *sListen) SetListenPort(u16Port uint16) error {
	if u16Port == 0 {
		return errors.New("error tcp port")
	}
	if this.m_Listener != nil {
		return errors.New("had listen")
	}
	this.m_u16Port = u16Port
	return nil
}

func (this *sListen) Listen() (err error) {
	if this.m_u16Port == 0 {
		return errors.New("error tcp port")
	}
	if this.m_Listener != nil {
		return errors.New("had listen")
	}
	if this.m_funcOnListen == nil {
		return errors.New("not set listen func")
	}

	strListenAddr := ":" + strconv.Itoa((int)(this.m_u16Port))
	GLog.Mustln("listen strListenAddr", strListenAddr)

	this.m_Listener, err = net.Listen("tcp", strListenAddr)
	if err != nil {
		return
	}

	go func() {
		defer bsn_common.FuncGuard()
		defer this.StopListen()

		for {
			conn, err := this.m_Listener.Accept()
			if err != nil {
				GLog.Errorln(err)
				return
			}

			this.m_funcOnListen(conn)
		}
	}()

	return
}

func (this *sListen) StopListen() {
	this.m_MutexStopListen.Lock()
	defer this.m_MutexStopListen.Unlock()

	if this.m_Listener == nil {
		return
	}
	this.m_Listener.Close()
	this.m_Listener = nil
}
