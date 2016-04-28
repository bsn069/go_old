package bsn_echo

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "unsafe"
	// "net"
	// "sync"
	"bsn_msg_client_echo_server"
	"bsn_msg_gate_server"
	"github.com/golang/protobuf/proto"
	"time"
)

type SClientUser struct {
	*bsn_net.SSessionWithMsgHeader
	*bsn_common.SState
	*bsn_common.SNotifyClose

	M_SClientUserMgr *SClientUserMgr
	M_TClientId      TClientId
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_TClientId:      0,
	}
	this.SSessionWithMsgHeader, _ = bsn_net.NewSSessionWithMsgHeader()
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
	go this.waitLogin()
	return nil
}

func (this *SClientUser) waitLogin() {
	time.Sleep(time.Duration(5) * time.Second)
	if this.Id() == 0 {
		GSLog.Debugln("wait login fail, close")
		this.Close()
	}
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
		GSLog.Debugln("close connect")
		this.Conn().Close()

		GSLog.Debugln("delete from user mgr")
		this.M_SClientUserMgr.delClient(this.Id())

		this.SNotifyClose.Close()
		this.Set(bsn_common.CState_Idle)
		GSLog.Debugln("close ok")
	}()

	this.Change(bsn_common.CState_Op, bsn_common.CState_Runing)
	for {
		err := this.RecvMsg()
		if err != nil {
			GSLog.Errorln(err)
			break
		}

		err = this.procMsg()
		if err != nil {
			GSLog.Errorln(err)
			break
		}
	}
}

func (this *SClientUser) Send2Client(vTClientId TClientId, byMsg []byte) error {
	vSMsg_Server2Gate_ClientMsg := new(bsn_msg.SMsg_Server2Gate_ClientMsg)
	vSMsg_Server2Gate_ClientMsg.Fill(uint16(vTClientId), byMsg)
	return this.SendMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_gate_server.ECmdServe2Gate_CmdServer2Gate_ClientMsg), vSMsg_Server2Gate_ClientMsg.Serialize())
}

func (this *SClientUser) SendPb2Client(vTClientId TClientId, msgType bsn_msg_client_echo_server.ECmdEchoServe2Client, iMessage proto.Message) error {
	byMsg, err := proto.Marshal(iMessage)
	if err != nil {
		GSLog.Errorln(err)
		return err
	}

	byData := bsn_msg.NewMsgWithMsgHeader(bsn_common.TMsgType(msgType), byMsg)
	return this.Send2Client(vTClientId, byData)
}
