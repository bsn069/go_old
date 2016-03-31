package bsn_net

// import (
// 	"errors"
// 	"github.com/bsn069/go/bsn_common"
// 	"net"
// 	// "strconv"
// 	"fmt"
// 	"sync"
// )

// type SListen struct {
// 	M_strAddr   string
// 	M_Listener  net.Listener
// 	M_Mutex     sync.Mutex
// 	M_chanClose bsn_common.TNetChanClose
// 	M_chanConn  bsn_common.TNetChanConn
// }

// func NewListen() *SListen {
// 	this := &SListen{
// 		M_chanClose: make(bsn_common.TNetChanClose, 0),
// 		M_chanConn:  make(bsn_common.TNetChanConn, 100),
// 	}
// 	return this
// }

// func (this *SListen) SetListenPort(u16Port uint16) error {
// 	this.M_Mutex.Lock()
// 	defer this.M_Mutex.Unlock()

// 	if u16Port < 1024 {
// 		return errors.New("error port must > 1024")
// 	}
// 	if this.M_Listener != nil {
// 		return errors.New("had listen")
// 	}
// 	this.M_strAddr = fmt.Sprintf(":%v", u16Port)
// 	return nil
// }

// func (this *SListen) SetListenAddr(strAddr string) error {
// 	this.M_Mutex.Lock()
// 	defer this.M_Mutex.Unlock()

// 	if strAddr == "" {
// 		return errors.New("error addres")
// 	}
// 	if this.M_Listener != nil {
// 		return errors.New("had listen")
// 	}
// 	this.M_strAddr = strAddr
// 	return nil
// }

// func (this *SListen) IsListen() bool {
// 	return this.M_Listener != nil
// }

// func (this *SListen) Listen() (err error) {
// 	this.M_Mutex.Lock()
// 	defer this.M_Mutex.Unlock()

// 	if this.M_strAddr == "" {
// 		return errors.New("no address")
// 	}
// 	if this.M_Listener != nil {
// 		return errors.New("had listen")
// 	}

// 	GSLog.Mustln("listen strListenAddr ", this.M_strAddr)
// 	this.M_Listener, err = net.Listen("tcp", this.M_strAddr)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	// if not call StopListen, clear close chanel
// 	select {
// 	case <-this.M_chanClose:
// 	default:
// 	}

// 	go this.listenFunc()
// 	return
// }

// // block until close
// func (this *SListen) StopListen() error {
// 	this.M_Mutex.Lock()
// 	defer this.M_Mutex.Unlock()
// 	// GSLog.Debugln("1")

// 	if this.M_Listener == nil {
// 		GSLog.Debugln("2")
// 		return errors.New("not listen")
// 	}
// 	// GSLog.Debugln("3")
// 	err := this.M_Listener.Close()
// 	if err != nil {
// 		GSLog.Debugln("4" + err.Error())
// 	}

// 	// close(this.M_chanConn)
// 	// GSLog.Debugln("5")
// 	<-this.M_chanClose
// 	// GSLog.Debugln("6")
// 	this.M_Listener = nil
// 	// GSLog.Debugln("7")
// 	return nil
// }

// func (this *SListen) listenFunc() {
// 	GSLog.Mustln("listenFunc")

// 	defer bsn_common.FuncGuard()
// 	defer func() {
// 		// GSLog.Debugln("send close before")
// 		this.M_chanClose <- true
// 		// GSLog.Debugln("send close after")
// 	}()

// 	for {
// 		// GSLog.Debugln("wait accept")
// 		vConn, err := this.M_Listener.Accept()
// 		// GSLog.Debugln("have accept")
// 		if err != nil {
// 			GSLog.Errorln(err)
// 			return
// 		}

// 		// GSLog.Debugln("send conn before")
// 		this.M_chanConn <- vConn
// 		// GSLog.Debugln("send conn after")
// 	}
// }
