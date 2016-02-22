package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
)

const (
	CSMsgHeader_Size bsn_common.TMsgLen = 4
)

type SMsgHeader struct {
	M_TMsgType bsn_common.TMsgType
	M_TMsgLen  bsn_common.TMsgLen
}

func NewMsgHeader(vTMsgType bsn_common.TMsgType, vTMsgLen bsn_common.TMsgLen) *SMsgHeader {
	this := &SMsgHeader{}
	this.Fill(vTMsgType, vTMsgLen)
	return this
}

func (this *SMsgHeader) Fill(vTMsgType bsn_common.TMsgType, vTMsgLen bsn_common.TMsgLen) {
	this.M_TMsgType = vTMsgType
	this.M_TMsgLen = vTMsgLen
}

func (this *SMsgHeader) Len() bsn_common.TMsgLen {
	return this.M_TMsgLen
}

func (this *SMsgHeader) Type() bsn_common.TMsgType {
	return this.M_TMsgType
}

func (this *SMsgHeader) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeader_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsgHeader) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	bsn_common.WriteUint16(byDatas, uint16(this.Type()))
	bsn_common.WriteUint16(byDatas[2:], uint16(this.Len()))
	return bsn_common.TMsgLen(4)
}

func (this *SMsgHeader) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	this.M_TMsgType = bsn_common.TMsgType(bsn_common.ReadUint16(byDatas))
	this.M_TMsgLen = bsn_common.TMsgLen(bsn_common.ReadUint16(byDatas[2:]))
	return bsn_common.TMsgLen(4)
}
