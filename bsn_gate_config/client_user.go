package bsn_gate_config

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
	M_SMsgHeader     *bsn_msg.SMsgHeader
	M_by2MsgBody     []byte
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_TClientId:      0,
		M_byRecvBuff:     make([]byte, 4),
		M_SMsgHeader:     new(bsn_msg.SMsgHeader),
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
	for {
		GSLog.Debugln("read msg header")
		byMsgHeader := this.M_byRecvBuff[0:bsn_msg.CSMsgHeader_Size]
		err := this.Recv(byMsgHeader)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
		GSLog.Debugln("recv byMsgHeader= ", byMsgHeader)

		this.M_SMsgHeader.DeSerialize(byMsgHeader)
		GSLog.Debugln("recv this.M_SMsgHeader= ", this.M_SMsgHeader)

		vTotalLen := int(this.M_SMsgHeader.Len())
		if vTotalLen > cap(this.M_byRecvBuff) {
			// realloc recv buffer
			this.M_byRecvBuff = make([]byte, vTotalLen)
		}

		GSLog.Debugln("read this.M_by2MsgBody")
		this.M_by2MsgBody = this.M_byRecvBuff[0:vTotalLen]
		err = this.Recv(this.M_by2MsgBody)
		if err != nil {
			GSLog.Errorln(err)
			break
		}
		GSLog.Debugln("recv this.M_by2MsgBody= ", this.M_by2MsgBody)

		err = this.procMsg()
		if err != nil {
			GSLog.Errorln(err)
			break
		}
	}
}

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)
	return nil
}
