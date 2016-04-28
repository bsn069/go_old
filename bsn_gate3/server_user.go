package bsn_gate3

import (
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	// "bsn_msg_gate_server"
	// "github.com/golang/protobuf/proto"
)

type SServerUser struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
	M_ServerType     uint8

	M_MsgTypeMin bsn_common.TMsgType
	M_MsgTypeMax bsn_common.TMsgType
}

func NewSServerUser(vSServerUserMgr *SServerUserMgr) (*SServerUser, error) {
	GSLog.Debugln("NewSServerUser()")

	this := &SServerUser{
		M_SServerUserMgr: vSServerUserMgr,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

	return this, nil
}

func (this *SServerUser) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUser) ServerType() uint8 {
	return this.M_ServerType
}

func (this *SServerUser) OnClientMsg(vSClientUser *SClientUser) bool {
	if vSClientUser.MsgType() < this.M_MsgTypeMin || vSClientUser.MsgType() > this.M_MsgTypeMax {
		return false
	}

	return this.send_CmdGate2Server_ClientMsg(vSClientUser)
}

func (this *SServerUser) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerUser) ShowInfo() {
}

func (this *SServerUser) Run() (err error) {
	err = this.SConnecterWithMsgHeader.Run()
	if err != nil {
		return
	}

	return this.send_CmdGate2Server_LoginReq()
}
