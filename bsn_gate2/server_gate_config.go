package bsn_gate2

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	"errors"
	// "net"
	// "strconv"
	// "sync"
	"fmt"
)

type SServerGateConfig struct {
	*bsn_net.SConnecterWithMsgHeader
	M_SGate *SGate
}

func NewSServerGateConfig(vSGate *SGate, strAddr string) (*SServerGateConfig, error) {
	GSLog.Debugln("NewSServerGateConfig()")

	this := &SServerGateConfig{
		M_SGate: vSGate,
	}
	this.SConnecterWithMsgHeader, _ = bsn_net.NewSConnecterWithMsgHeader(this)
	this.SetAddr(strAddr)

	return this, nil
}

func (this *SServerGateConfig) NetConnecterWithMsgHeaderImpOnClose() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpOnClose")
	return nil
}

func (this *SServerGateConfig) App() *SGate {
	return this.M_SGate
}

func (this *SServerGateConfig) ShowInfo() {
}

func (this *SServerGateConfig) Run() {
	this.SConnecterWithMsgHeader.Run()
	this.SendMsgWithSMsgHeader(bsn_msg.GMsgDefine_Gate2GateConfig_GateReg, nil)
}

func (this *SServerGateConfig) NetConnecterWithMsgHeaderImpProcMsg() error {
	GSLog.Debugln("NetConnecterWithMsgHeaderImpProcMsg")

	switch this.M_SMsgHeader.Type() {
	case bsn_msg.GMsgDefine_Gate2GateConfig_GateReg:
		return this.ProcMsg_GateReg()
	}

	strInfo := fmt.Sprintf("nuknown msg type=%u", this.M_SMsgHeader.Type())
	return errors.New(strInfo)
}

func (this *SServerGateConfig) ProcMsg_GateReg() error {
	GSLog.Debugln("ProcMsg_GateReg")
	this.App().M_chanWaitGateConfig <- true
	return nil
}
