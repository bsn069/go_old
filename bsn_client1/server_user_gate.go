package bsn_client1

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	// "fmt"
)

type SServerUserGate struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
}

func NewSServerUserGate(vSServerUserMgr *SServerUserMgr, strAddr string) (*SServerUserGate, error) {
	GSLog.Debugln("NewSServerUserGate()")

	this := &SServerUserGate{
		M_SServerUserMgr: vSServerUserMgr,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)
	this.SetAddr(strAddr)

	return this, nil
}
func (this *SServerUserGate) UserMgr() *SServerUserMgr {
	return this.M_SServerUserMgr
}

func (this *SServerUserGate) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerUserGate) ShowInfo() {
}

func (this *SServerUserGate) Run() {
	this.SConnecterWithMsgHeader.Run()
}

func (this *SServerUserGate) Close() {
	this.SConnecterWithMsgHeader.Close()
}

func (this *SServerUserGate) GatePing() error {
	return this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Client2Gate_Ping, nil)
}

func (this *SServerUserGate) GateEcho(strInfo string) error {
	return this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Client2Gate_Echo, []byte(strInfo))
}

func (this *SServerUserGate) EchoPing() error {
	return this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Client2Echo_Ping, nil)
}

func (this *SServerUserGate) EchoEcho(strInfo string) error {
	return this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Client2Echo_Echo, []byte(strInfo))
}
