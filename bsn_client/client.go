package bsn_client

import (
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	"strconv"
	// "sync"
)

type TClientId uint16

var GClientId TClientId = 0

type SClient struct {
	bsn_net.SNetConnecter
	M_TClientId  TClientId
	M_SMsgHeader bsn_msg.SMsgHeader
}

func NewClient() (*SClient, error) {
	GClientId++
	GSLog.Debugln("NewClient() GClientId=", GClientId)

	this := &SClient{
		M_TClientId: GClientId,
	}
	this.SNetConnecter.M_INetConnecterImp = this

	vSCmdClient := &SCmdClient{M_SClient: this}
	bsn_input.GSInput.Reg("Client"+strconv.Itoa(int(GClientId)), vSCmdClient)

	return this, nil
}

func (this *SClient) ShowInfo() {
	GSLog.Mustln("id :", this.M_TClientId)
}

func (this *SClient) NetConnecterImpRun() error {
	GSLog.Debugln("NetConnecterImpRun")
	vTMsgType, byData, err := this.RecvMsgWithSMsgHeader()
	if err != nil {
		if err.Error() == "EOF" {
			GSLog.Errorln("connect disconnect")
		} else {
			GSLog.Errorln("ReadMsg error: ", err)
		}
		return err
	}
	GSLog.Debugln("recv msg: ", vTMsgType, byData)

	return err
}

func (this *SClient) NetConnecterImpOnClose() error {
	GSLog.Debugln("NetConnecterImpOnClose")
	return nil
}

func (this *SClient) SendString(strInfo string) error {
	this.SendMsgWithSMsgHeader(0, []byte(strInfo))
	return nil
}
