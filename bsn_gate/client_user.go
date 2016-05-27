package bsn_gate

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	// "fmt"
	"net"
	// "sync/atomic"
	"sync/atomic"
)

type SClientUser struct {
	M_SClientUserMgr *SClientUserMgr
	M_TClientId      uint16
	M_Recver         bsn_net.STCPRecver
	M_CloseReason    int32
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr, vConn net.Conn, vClientId uint16) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_TClientId:      vClientId,
	}

	this.M_Recver.Reset()
	this.M_Recver.SetIRecverCB(this)
	this.M_Recver.SetConn(vConn)

	return this, nil
}

func (this *SClientUser) RecverCBOnMsg(vSMsg *bsn_msg.SMsg) (err error) {
	GSLog.Debugln("RecverCBOnMsg ", vSMsg.Type())
	err = this.procMsg(vSMsg)
	vSMsg.Del()
	return err
}

func (this *SClientUser) RecverCBOnClose() (err error) {
	GSLog.Debugln("RecverCBOnClose ", this.ClientId())
	this.setCloseReason(bsn_common.CCloseReason_Disconnect)
	// save data
	GSLog.Debugln("this.M_CloseReason= ", this.M_CloseReason)
	return nil
}

func (this *SClientUser) ClientId() uint16 {
	return this.M_TClientId
}

// CCloseReason_Idle
func (this *SClientUser) setCloseReason(vCCloseReason int32) bool {
	return atomic.CompareAndSwapInt32((*int32)(&this.M_CloseReason), bsn_common.CCloseReason_Idle, vCCloseReason)
}

func (this *SClientUser) run() (err error) {
	err = this.M_Recver.Start()
	return err
}

func (this *SClientUser) close() (err error) {
	GSLog.Debugln("close ", this.ClientId())

	this.M_Recver.Stop()

	return nil
}

func (this *SClientUser) procMsg(vMsg *bsn_msg.SMsg) (err error) {
	return nil
}
