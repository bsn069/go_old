package bsn_net

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
)

type INetConnecterWithMsgHeaderImp interface {
	NetConnecterWithMsgHeaderImpProcMsg() error
	NetConnecterWithMsgHeaderImpOnClose() error
}

type SConnecterWithMsgHeader struct {
	*SNetConnecter

	M_byRecvBuff []byte
	M_SMsgHeader *bsn_msg.SMsgHeader
	M_by2MsgBody []byte

	M_INetConnecterWithMsgHeaderImp INetConnecterWithMsgHeaderImp
}

func NewSConnecterWithMsgHeader(vINetConnecterWithMsgHeaderImp INetConnecterWithMsgHeaderImp) (*SConnecterWithMsgHeader, error) {
	GSLog.Debugln("NewSConnecterWithMsgHeader()")

	this := &SConnecterWithMsgHeader{
		M_byRecvBuff:                    make([]byte, 4),
		M_SMsgHeader:                    new(bsn_msg.SMsgHeader),
		M_INetConnecterWithMsgHeaderImp: vINetConnecterWithMsgHeaderImp,
	}
	this.SNetConnecter, _ = NewNetConnecter(this)

	return this, nil
}

func (this *SConnecterWithMsgHeader) NetConnecterImpRun() error {
	GSLog.Debugln("NetConnecterImpRun")

	GSLog.Debugln("read msg header")
	byMsgHeader := this.M_byRecvBuff[0:bsn_msg.CSMsgHeader_Size]
	err := this.Recv(byMsgHeader)
	if err != nil {
		GSLog.Errorln(err)
		return err
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
	if vTotalLen > 0 {
		err = this.Recv(this.M_by2MsgBody)
		if err != nil {
			GSLog.Errorln(err)
			return err
		}
	}
	GSLog.Debugln("recv this.M_by2MsgBody= ", this.M_by2MsgBody)

	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)
	err = this.M_INetConnecterWithMsgHeaderImp.NetConnecterWithMsgHeaderImpProcMsg()
	if err != nil {
		GSLog.Errorln(err)
		return err
	}

	return nil
}

func (this *SConnecterWithMsgHeader) NetConnecterImpOnClose() error {
	GSLog.Debugln("NetConnecterImpOnClose")
	return this.M_INetConnecterWithMsgHeaderImp.NetConnecterWithMsgHeaderImpOnClose()
}
