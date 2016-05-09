package bsn_gate

import (
	"errors"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	// "fmt"
	"net"
	// "sync"
)

type SClientUser struct {
	M_SClientUserMgr *SClientUserMgr
	M_Conn           net.Conn
	M_TClientId      uint16
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr, vConn net.Conn, vClientId uint16) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_Conn:           vConn,
		M_TClientId:      vClientId,
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

func (this *SClientUser) close() (err error) {
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
	for {
		readLen, err := vConn.Read(byMsgHeader)
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

	if err != nil {
		GSLog.Debugln(err)
	}

	vConn.Close()
	this.M_Conn = nil

	this.M_SClientUserMgr.delClient(this.ClientId())
	return
}
