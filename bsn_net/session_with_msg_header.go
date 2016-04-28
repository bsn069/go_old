package bsn_net

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "net"
)

type SSessionWithMsgHeader struct {
	*SSessionAddr

	M_byRecvBuff []byte
	M_SMsgHeader *bsn_msg.SMsgHeader
	M_by2MsgBody []byte
}

func NewSSessionWithMsgHeader() (*SSessionWithMsgHeader, error) {
	GSLog.Debugln("NewSSession")

	this := &SSessionWithMsgHeader{
		M_byRecvBuff: make([]byte, 4),
		M_SMsgHeader: new(bsn_msg.SMsgHeader),
	}
	this.SSessionAddr, _ = NewSSessionAddr()

	return this, nil
}

func (this *SSessionWithMsgHeader) MsgType() bsn_common.TMsgType {
	return this.M_SMsgHeader.Type()
}

func (this *SSessionWithMsgHeader) RecvMsg() error {
	GSLog.Debugln("RecvMsg")

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

	return nil
}
