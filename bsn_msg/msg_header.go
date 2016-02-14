package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
)

type SMsgHeader struct {
	M_TMsgType TMsgType
	M_TMsgLen  TMsgLen
}

func newMsgHeader(vTMsgType TMsgType, vTMsgLen TMsgLen) *SMsgHeader {
	this := &SMsgHeader{}
	this.Fill(vTMsgType, vTMsgLen)
	return this
}

func (this *SMsgHeader) Fill(vTMsgType TMsgType, vTMsgLen TMsgLen) {
	this.M_TMsgType = vTMsgType
	this.M_TMsgLen = vTMsgLen
}

func (this *SMsgHeader) Len() TMsgLen {
	return this.M_TMsgLen
}

func (this *SMsgHeader) Type() TMsgType {
	return this.M_TMsgType
}

func (this *SMsgHeader) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeader_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsgHeader) Serialize2Byte(byDatas []byte) TMsgLen {
	bsn_common.WriteUint16(byDatas, uint16(this.Type()))
	bsn_common.WriteUint16(byDatas[2:], uint16(this.Len()))
	return TMsgLen(4)
}

func (this *SMsgHeader) DeSerialize(byDatas []byte) TMsgLen {
	this.M_TMsgType = TMsgType(bsn_common.ReadUint16(byDatas))
	this.M_TMsgLen = TMsgLen(bsn_common.ReadUint16(byDatas[2:]))
	return TMsgLen(4)
}
