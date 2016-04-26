package bsn_client

import (
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	"errors"
	// "net"
	"strconv"
	// "sync"
	"fmt"
)

type TClientId uint16

var GClientId TClientId = 0

type SClient struct {
	*bsn_net.SConnecterWithMsgHeader
	M_TClientId TClientId
}

func NewClient(strAddr string) (*SClient, error) {
	GClientId++
	GSLog.Debugln("NewClient() GClientId=", GClientId)

	this := &SClient{
		M_TClientId: GClientId,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)
	this.SetAddr(strAddr)

	vSCmdClient := &SCmdClient{M_SClient: this}
	bsn_input.GSInput.Reg("Client"+strconv.Itoa(int(GClientId)), vSCmdClient)

	return this, nil
}

func (this *SClient) ShowInfo() {
	GSLog.Mustln("id :", this.M_TClientId)
}

func (this *SClient) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SClient) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")

	switch this.M_SMsgHeader.Type() {
	case bsn_msg.GMsgDefine_Echo2Client_Test:
		return this.ProcMsg_Test()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", this.M_SMsgHeader.Type())
	return errors.New(strInfo)
}

func (this *SClient) ProcMsg_Test() error {
	GSLog.Debugln("ProcMsg_Test")
	return nil
}

func (this *SClient) SendString(strInfo string) error {
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Client2Echo_Test, []byte(strInfo))
	return nil
}
