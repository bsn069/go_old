package bsn_net

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"net"
	"strconv"
	"sync"
)

type sListen struct {
	m_port TPort

	m_iListener       net.Listener
	m_Mutex           sync.Mutex
	m_chanClose       TCloseChan
	m_iListenCallBack IListenCallBack
}

func newListen(iListenCallBack IListenCallBack) IListen {
	this := &sListen{
		m_iListenCallBack: iListenCallBack,
		m_chanClose:       make(TCloseChan, 1),
	}
	return this
}

func (this *sListen) SetListenPort(port TPort) error {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if port == 0 {
		return errors.New("error tcp port")
	}
	if this.m_iListener != nil {
		return errors.New("had listen")
	}
	this.m_port = port
	return nil
}

func (this *sListen) Listen() (err error) {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	if this.m_port == 0 {
		return errors.New("error tcp port")
	}
	if this.m_iListener != nil {
		return errors.New("had listen")
	}
	if this.m_iListenCallBack == nil {
		return errors.New("not set IListenCallBack")
	}

	strListenAddr := ":" + strconv.Itoa((int)(this.m_port))
	GLog.Mustln("listen strListenAddr", strListenAddr)

	this.m_iListener, err = net.Listen("tcp", strListenAddr)
	if err != nil {
		return
	}

	select {
	case <-this.m_chanClose:
	default:
	}

	go this.listenFunc()
}

// block until close
func (this *sListen) StopListen() {
	this.m_Mutex.Lock()
	defer this.m_Mutex.Unlock()

	iLisnter := this.m_iListener
	if this.m_iListener == nil {
		return
	}
	iLisnter.Close()
	<-this.m_chanClose
}

func (this *sListen) listenFunc() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.m_iListener.Close()
		this.m_iListener = nil
		this.m_chanClose <- true
	}()

	for {
		iConn, err := this.m_iListener.Accept()
		if err != nil {
			GLog.Errorln(err)
			return
		}

		err = this.m_iListenCallBack.OnListen(iConn)
		if err != nil {
			GLog.Errorln(err)
			iConn.Close()
		}
	}
}
