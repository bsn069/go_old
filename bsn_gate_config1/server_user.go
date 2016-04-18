package bsn_gate3

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
)

type SServerUser struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
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

func (this *SServerUser) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerUser) ShowInfo() {
}

func (this *SServerUser) Ping(strInfo string) error {
	// return this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2Server_Ping, []byte(strInfo))
	return nil
}
