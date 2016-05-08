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
	"bsn_define"
	"time"
)

type SServerUser struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
	M_ServerType     bsn_define.EServerType

	M_MsgTypeMin bsn_common.TMsgType
	M_MsgTypeMax bsn_common.TMsgType

	M_Close bool
}

func NewSServerUser(vSServerUserMgr *SServerUserMgr) (*SServerUser, error) {
	GSLog.Debugln("NewSServerUser()")

	this := &SServerUser{
		M_SServerUserMgr: vSServerUserMgr,
		M_ServerType:     bsn_define.EServerType_ServerType_Counts,
		M_Close:          false,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

	return this, nil
}

func (this *SServerUser) HadLogin() bool {
	return this.M_ServerType != bsn_define.EServerType_ServerType_Counts
}

func (this *SServerUser) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUser) ServerType() bsn_define.EServerType {
	return this.M_ServerType
}

func (this *SServerUser) SetServerType(vEServerType bsn_define.EServerType) {
	this.M_ServerType = vEServerType
}

func (this *SServerUser) OnClientMsg(vSClientUser *SClientUser) bool {
	if vSClientUser.MsgType() < this.M_MsgTypeMin || vSClientUser.MsgType() > this.M_MsgTypeMax {
		return false
	}

	return this.send_CmdGate2Server_ClientMsg(vSClientUser)
}

func (this *SServerUser) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose", this.M_Close)
	if !this.M_Close {
		GSLog.Debugln("reconnect")
		go func() {
			for {
				time.Sleep(time.Duration(3) * time.Second)
				err := this.Run()
				if err == nil {
					break
				}
				GSLog.Debugln(err)
			}
		}()
	}
	return nil
}

func (this *SServerUser) ShowInfo() {
}
