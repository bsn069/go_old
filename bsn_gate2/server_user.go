package bsn_gate2

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_log"
	"errors"
	"net"
	// "strconv"
	"sync"
)

type SServerUser struct {
	M_SServerUserMgr          *SServerUserMgr
	M_SClientUserMgr          *SClientUserMgr
	M_strAddr                 string
	M_Conn                    net.Conn
	M_bRun                    bool
	M_Mutex                   sync.Mutex
	M_chanNotifyClose         chan bool
	M_chanWaitClose           chan bool
	M_byRecvBuff              []byte
	M_bySMsgHeaderServer2Gate []byte
	M_by2GateMsg              []byte
	M_by2ClientMsg            []byte
	M_SMsgHeaderServer2Gate   bsn_msg.SMsgHeaderServer2Gate
}

func NewSServerUser(vSServerUserMgr *SServerUserMgr, strAddr string) (*SServerUser, error) {
	GSLog.Debugln("NewSServerUser()")

	this := &SServerUser{
		M_SServerUserMgr:          vSServerUserMgr,
		M_SClientUserMgr:          vSServerUserMgr.Gate().GetClientMgr(),
		M_strAddr:                 strAddr,
		M_chanNotifyClose:         make(chan bool, 1),
		M_chanWaitClose:           make(chan bool, 1),
		M_bySMsgHeaderServer2Gate: make([]byte, bsn_msg.CSMsgHeaderServe2Gater_Size),
		M_byRecvBuff:              make([]byte, 4),
	}

	return this, nil
}

func (this *SServerUser) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUser) ShowInfo() {
}

func (this *SServerUser) Run() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if this.M_bRun {
		return errors.New("running")
	}

	err := this.Connect()
	if err != nil {
		return err
	}

	go this.runImp()
	this.M_bRun = true
	return nil
}

func (this *SServerUser) Close() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if !this.M_bRun {
		return errors.New("not running")
	}
	GSLog.Mustln("Close begin")

	this.Conn().Close()
	// clear close chanel
	select {
	case <-this.M_chanNotifyClose:
	default:
	}
	this.M_chanNotifyClose <- true
	// wait close complete
	<-this.M_chanWaitClose

	GSLog.Mustln("Close end")
	return nil
}

func (this *SServerUser) Connect() (err error) {
	if "" == this.M_strAddr {
		return errors.New("no addr")
	}

	this.M_Conn, err = net.Dial("tcp", this.M_strAddr)
	if err != nil {
		return err
	}

	return nil
}

func (this *SServerUser) Conn() net.Conn {
	return this.M_Conn
}

func (this *SServerUser) Send(vbyMsg []byte) error {
	vLen := len(vbyMsg)
	if vLen <= 0 {
		return nil
	}

	writeLen, err := this.Conn().Write(vbyMsg)
	if err != nil {
		return err
	}
	if writeLen != vLen {
		return errors.New("not write all data")
	}
	return nil
}

func (this *SServerUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("on closing")
		this.M_bRun = false

		GSLog.Debugln("close connect")
		this.Conn().Close()

		GSLog.Debugln("send notify to wait close chan")
		select {
		case <-this.M_chanWaitClose:
		default:
		}
		this.M_chanWaitClose <- true

		GSLog.Debugln("close ok")
	}()

	for {
		GSLog.Debugln("read msg header")
		err := this.readMsg(this.M_bySMsgHeaderServer2Gate)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
		GSLog.Debugln("recv this.M_bySMsgHeaderServer2Gate= ", this.M_bySMsgHeaderServer2Gate)

		this.M_SMsgHeaderServer2Gate.DeSerialize(this.M_bySMsgHeaderServer2Gate)
		GSLog.Debugln("recv this.M_SMsgHeaderServer2Gate= ", this.M_SMsgHeaderServer2Gate)

		vTotalLen := int(this.M_SMsgHeaderServer2Gate.Len()) + int(this.M_SMsgHeaderServer2Gate.ServerMsgLen())
		if vTotalLen > cap(this.M_byRecvBuff) {
			// realloc recv buffer
			this.M_byRecvBuff = make([]byte, vTotalLen)
		}
		this.M_byRecvBuff = this.M_byRecvBuff[0:vTotalLen]

		GSLog.Debugln("read byMsgBody")
		err = this.readMsg(this.M_byRecvBuff)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
		GSLog.Debugln("recv this.M_byRecvBuff= ", this.M_byRecvBuff)

		this.M_by2GateMsg = this.M_byRecvBuff[0:this.M_SMsgHeaderServer2Gate.Len()]
		this.M_by2ClientMsg = this.M_byRecvBuff[this.M_SMsgHeaderServer2Gate.Len():vTotalLen]
		err = this.procMsg()
		if err != nil {
			GSLog.Errorln(err)
			break
		}
	}
}

func (this *SServerUser) procMsg() error {
	GSLog.Debugln("this.M_SMsgHeaderServer2Gate= ", this.M_SMsgHeaderServer2Gate)
	GSLog.Debugln("this.M_by2GateMsg= ", this.M_by2GateMsg)
	GSLog.Debugln("this.M_by2ClientMsg= ", this.M_by2ClientMsg)
	return nil
}

func (this *SServerUser) readMsg(byMsg []byte) error {
	vLen := len(byMsg)
	if vLen <= 0 {
		return nil
	}

	readLen, err := this.Conn().Read(byMsg)
	if err != nil {
		return err
	}
	if readLen != vLen {
		return errors.New("not read all data")
	}
	return nil
}
