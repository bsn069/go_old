package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	"encoding/binary"
	"net"
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

func NewMsgHeaderFromBytes(byDatas []byte) *SMsgHeader {
	this := &SMsgHeader{}
	this.DeSerialize(byDatas)
	return this
}

func ReadMsgWithMsgHeader(conn net.Conn) (err error, u16Type uint16, byData []byte) {
	byMsgHeader := make([]byte, CSMsgHeader_Size)
	_, err = conn.Read(byMsgHeader)
	if err != nil {
		return
	}
	vSMsgHeader := NewMsgHeaderFromBytes(byMsgHeader)
	u16Type = uint16(vSMsgHeader.Type())

	vu16Len := uint16(vSMsgHeader.Len())
	if vu16Len > 0 {
		byData = make([]byte, vu16Len)
		_, err = conn.Read(byData)
		if err != nil {
			return
		}
	}

	return
}

func WriteMsgWithMsgHeader(conn net.Conn, u16Type uint16, byData []byte) (err error) {
	vTMsgLen := bsn_common.TMsgLen(len(byData))
	vSMsgHeader := NewMsgHeader(bsn_common.TMsgType(u16Type), vTMsgLen)

	_, err = conn.Write(vSMsgHeader.Serialize())
	if err != nil {
		return
	}

	if vTMsgLen > 0 {
		_, err = conn.Write(byData)
		if err != nil {
			return
		}
	}

	return
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
	var vTMsgLen bsn_common.TMsgLen = 0

	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], uint16(this.Type()))
	vTMsgLen += 2
	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], uint16(this.Len()))
	vTMsgLen += 2

	return vTMsgLen
}

func (this *SMsgHeader) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	var vTMsgLen bsn_common.TMsgLen = 0

	this.M_TMsgType = bsn_common.TMsgType(binary.LittleEndian.Uint16(byDatas[vTMsgLen:]))
	vTMsgLen += 2
	this.M_TMsgLen = bsn_common.TMsgLen(binary.LittleEndian.Uint16(byDatas[vTMsgLen:]))
	vTMsgLen += 2

	return vTMsgLen
}
