package bsn_client

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	"net"
	// "strconv"
	// "sync"
	// "bsn_msg_gate_server"
	// "github.com/golang/protobuf/proto"
	// "bsn_define"
	// "time"
)

type SGate struct {
	M_SApp       *SApp
	M_TCPstrAddr string
	M_Conn       net.Conn
}

func NewSGate(vSApp *SApp) (this *SGate, err error) {
	GSLog.Debugln("NewSGate")
	this = &SGate{
		M_SApp:       vSApp,
		M_TCPstrAddr: "127.0.0.1:20001",
	}

	return
}

func (this *SGate) start() (err error) {
	for i := 0; i < 10; i++ {
		err = this.connect()
		if err == nil {
			break
		}
	}

	return
}

func (this *SGate) stop() (err error) {
	if this.M_Conn != nil {
		this.M_Conn.Close()
		this.M_Conn = nil
	}
	return
}

func (this *SGate) connect() (err error) {
	GSLog.Mustln("connect gate")
	if this.M_Conn != nil {
		return
	}

	this.M_Conn, err = net.Dial("tcp", this.M_TCPstrAddr)
	if err != nil {
		return
	}

	GSLog.Mustln("connect gate success")
	return
}
