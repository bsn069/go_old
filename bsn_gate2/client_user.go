package bsn_gate2

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "unsafe"
	"net"
	"sync"
)

type TClientId uint16

type SClientUser struct {
	M_SClientUserMgr  *SClientUserMgr
	M_TClientId       TClientId
	M_Conn            net.Conn
	M_byRecvBuff      []byte
	M_Mutex           sync.Mutex
	M_bRun            bool
	M_chanWaitClose   chan bool
	M_chanNotifyClose chan bool
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr:  vSClientUserMgr,
		M_TClientId:       0,
		M_Conn:            nil,
		M_byRecvBuff:      make([]byte, 4),
		M_chanNotifyClose: make(chan bool, 1),
		M_chanWaitClose:   make(chan bool, 1),
	}

	return this, nil
}

func (this *SClientUser) UserMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SClientUser) SetId(vTClientId TClientId) error {
	this.M_TClientId = vTClientId
	return nil
}

func (this *SClientUser) Id() TClientId {
	return this.M_TClientId
}

func (this *SClientUser) SetConn(vConn net.Conn) error {
	this.M_Conn = vConn
	return nil
}

func (this *SClientUser) Conn() net.Conn {
	return this.M_Conn
}

func (this *SClientUser) Run() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if this.M_bRun {
		return errors.New("running")
	}

	go this.runImp()
	this.M_bRun = true
	return nil
}

func (this *SClientUser) Close() error {
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

	return nil
}

func (this *SClientUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("on closing")
		this.M_bRun = false

		GSLog.Debugln("close connect")
		this.Conn().Close()
		this.SetConn(nil)

		GSLog.Debugln("close from user mgr")
		this.M_SClientUserMgr.delClient(this.Id())
		this.M_SClientUserMgr = nil

		GSLog.Debugln("send notify to wait close chan")
		select {
		case <-this.M_chanWaitClose:
		default:
		}
		this.M_chanWaitClose <- true

		GSLog.Debugln("close ok")
	}()

	vSServerUserMgr := this.UserMgr().Gate().GetServerMgr()
	vMsg := new(bsn_msg.SMsgHeader)
	for {
		GSLog.Debugln("read msg header")
		byMsgHeader := this.M_byRecvBuff[0:bsn_msg.CSMsgHeader_Size]
		err := this.readMsg(byMsgHeader)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
		GSLog.Debugln("recv byMsgHeader= ", byMsgHeader)

		vMsg.DeSerialize(byMsgHeader)
		GSLog.Debugln("recv vMsg= ", vMsg)

		vTotalLen := int(vMsg.Len() + bsn_msg.CSMsgHeader_Size)
		if vTotalLen > cap(this.M_byRecvBuff) {
			// realloc recv buffer
			this.M_byRecvBuff = make([]byte, vTotalLen)
			copy(this.M_byRecvBuff, byMsgHeader)
		}

		GSLog.Debugln("read byMsgBody")
		byMsgBody := this.M_byRecvBuff[bsn_msg.CSMsgHeader_Size:vTotalLen]
		err = this.readMsg(byMsgBody)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
		GSLog.Debugln("recv byMsgBody= ", byMsgBody)

		byMsg := this.M_byRecvBuff[0:vTotalLen]
		GSLog.Debugln("recv byMsg= ", byMsg)
		err = vSServerUserMgr.Send(this, vMsg, byMsg)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
	}
}

func (this *SClientUser) readMsg(byMsg []byte) error {
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
