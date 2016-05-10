package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	// "fmt"
	"net"
	// "sync/atomic"
)

type SClientUser struct {
	M_SClientUserMgr    *SClientUserMgr
	M_Conn              net.Conn
	M_TClientId         uint16
	M_SStateCloseReason *bsn_common.SState // 0 1conenct disconnect 1usermgr close
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr, vConn net.Conn, vClientId uint16) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr:    vSClientUserMgr,
		M_Conn:              vConn,
		M_TClientId:         vClientId,
		M_SStateCloseReason: bsn_common.NewSState(),
	}

	return this, nil
}

func (this *SClientUser) ClientId() uint16 {
	return this.M_TClientId
}

func (this *SClientUser) conn() net.Conn {
	return this.M_Conn
}

func (this *SClientUser) run() (err error) {
	go this.workerRecvMsg()
	return
}

func (this *SClientUser) setCloseReason(vState bsn_common.TState) bool {
	return this.M_SStateCloseReason.Change(bsn_common.CState_Idle, vState)
}

func (this *SClientUser) close() (err error) {
	GSLog.Debugln("close ", this.ClientId())
	vConn := this.conn()
	if vConn == nil {
		err = errors.New("not run")
		return
	}
	vConn.Close()
	return
}

func (this *SClientUser) procMsg(vMsg *bsn_msg.SMsg) (err error) {
	return
}

func (this *SClientUser) workerRecvMsg() (err error) {
	GSLog.Debugln("workerRecvMsg")

	byMsgHeader := make([]byte, bsn_msg.CSMsgHeader_Size)
	vConn := this.conn()
	readLen := 0
	for {
		readLen, err = vConn.Read(byMsgHeader)
		if err != nil {
			break
		}
		if readLen != int(bsn_msg.CSMsgHeader_Size) {
			err = errors.New("not read all header data")
			break
		}

		vSMsg := bsn_msg.NewSMsg()
		byMsgBody := vSMsg.MsgBodyBuffer(byMsgHeader)

		readLen, err = vConn.Read(byMsgBody)
		if err != nil {
			break
		}
		if readLen != int(vSMsg.M_SMsgHeader.Len()) {
			err = errors.New("not read all body data")
			break
		}

		err = this.procMsg(vSMsg)
		if err != nil {
			break
		}
	}
	this.setCloseReason(bsn_common.CState_CloseReasonDisconnect)

	if err != nil {
		GSLog.Debugln("err=", err)
	}

	vConn.Close()
	this.M_Conn = nil

	this.M_SClientUserMgr.delClient(this.ClientId())
	GSLog.Debugln("workerRecvMsg end", this.M_SStateCloseReason)
	return
}
