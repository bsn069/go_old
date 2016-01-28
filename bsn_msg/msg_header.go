package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
)

type sMsgHeader struct {
	m_u16Type uint16
	m_u16Len  uint16
}

func newMsgHeader(u16Type, u16Len uint16) IMsgHeader {
	return &sMsgHeader{m_u16Type: u16Type, m_u16Len: u16Len}
}

func (this *sMsgHeader) Len() uint16 {
	return this.m_u16Len
}

func (this *sMsgHeader) Type() uint16 {
	return this.m_u16Type
}

func (this *sMsgHeader) Serialize() []byte {
	var byDatas = make([]byte, CMsgHeader_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *sMsgHeader) Serialize2Byte(byDatas []byte) {
	bsn_common.WriteUint16(byDatas, this.Type())
	bsn_common.WriteUint16(byDatas[2:], this.Len())
}

func (this *sMsgHeader) DeSerialize(byDatas []byte) {
	this.m_u16Type = bsn_common.ReadUint16(byDatas)
	this.m_u16Len = bsn_common.ReadUint16(byDatas[2:])
}
