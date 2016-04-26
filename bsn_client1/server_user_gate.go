package bsn_client1

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
	// "fmt"
	"bsn_msg_client_echo_server"
	"github.com/golang/protobuf/proto"
)

type SServerUserGate struct {
	*bsn_net.SConnecterWithMsgHeader

	M_SServerUserMgr *SServerUserMgr
}

func NewSServerUserGate(vSServerUserMgr *SServerUserMgr) (*SServerUserGate, error) {
	GSLog.Debugln("NewSServerUserGate()")

	this := &SServerUserGate{
		M_SServerUserMgr: vSServerUserMgr,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)

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

func (this *SServerUserGate) TestReq() {
	sendMsg := &bsn_msg_client_echo_server.STestReq{
		VstrInfo: proto.String("client test req"),
	}
	this.SendPbMsgWithSMsgHeader(bsn_common.TMsgType(bsn_msg_client_echo_server.ECmdClient2EchoServer_CmdClient2EchoServer_TestReq), sendMsg)
}
