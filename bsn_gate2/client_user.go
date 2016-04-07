package bsn_gate2

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "unsafe"
	// "net"
	// "sync"
)

type TClientId uint16

type SClientUser struct {
	*bsn_net.SSession
	*bsn_common.SState
	*bsn_common.SNotifyClose

	M_SClientUserMgr *SClientUserMgr
	M_TClientId      TClientId
	M_byRecvBuff     []byte
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_TClientId:      0,
		M_byRecvBuff:     make([]byte, 4),
	}
	this.SSession, _ = bsn_net.NewSSession()
	this.SState = bsn_common.NewSState()
	this.SNotifyClose = bsn_common.NewSNotifyClose()

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

func (this *SClientUser) Run() (err error) {
	if !this.Change(bsn_common.CState_Idle, bsn_common.CState_Op) {
		return errors.New("had listen")
	}

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Idle)
	}()

	this.SNotifyClose.Clear()

	go this.runImp()
	return nil
}

func (this *SClientUser) Close() (err error) {
	if !this.Change(bsn_common.CState_Runing, bsn_common.CState_Op) {
		return errors.New("not listen")
	}
	GSLog.Debugln("close")

	defer func() {
		if err == nil {
			return
		}
		this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	}()

	this.Conn().Close()
	// this.SNotifyClose.NotifyClose()
	this.SNotifyClose.WaitClose()

	return nil
}

func (this *SClientUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("on closing")

		GSLog.Debugln("close connect")
		this.Conn().Close()

		GSLog.Debugln("close from user mgr")
		this.M_SClientUserMgr.delClient(this.Id())

		this.SNotifyClose.Close()
		this.Set(bsn_common.CState_Idle)
		GSLog.Debugln("close ok")
	}()

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	vSServerUserMgr := this.UserMgr().Gate().GetServerMgr()
	vMsg := new(bsn_msg.SMsgHeader)
	for {
		GSLog.Debugln("read msg header")
		byMsgHeader := this.M_byRecvBuff[0:bsn_msg.CSMsgHeader_Size]
		err := this.Recv(byMsgHeader)
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
		err = this.Recv(byMsgBody)
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
