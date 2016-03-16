package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	"encoding/binary"
)

type TServer2GateMsgType uint16

const (
	CServer2GateMsgType_Test TServer2GateMsgType = iota + 1
	CServer2GateMsgType_1

	CSMsgHeaderServe2Gater_Size bsn_common.TMsgLen = CSMsgHeader_Size + 2
)

// msgHeader from server
type SMsgHeaderServer2Gate struct {
	SMsgHeader
	M_TMsgLenServerMsg bsn_common.TMsgLen
}

func NewMsgHeaderServer2GateFromBytes(byDatas []byte) *SMsgHeaderServer2Gate {
	this := &SMsgHeaderServer2Gate{}
	this.DeSerialize(byDatas)
	return this
}

func (this *SMsgHeaderServer2Gate) FillMsg(vTServer2GateMsgType TServer2GateMsgType, by2GateMsg, by2ClientMsg []byte) {
	this.SMsgHeader.Fill(bsn_common.TMsgType(vTServer2GateMsgType), bsn_common.TMsgLen(len(by2GateMsg)))
	this.Fill(bsn_common.TMsgLen(len(by2ClientMsg)))
}

func (this *SMsgHeaderServer2Gate) Fill(vTMsgLenServerMsg bsn_common.TMsgLen) {
	this.M_TMsgLenServerMsg = vTMsgLenServerMsg
}

// send to client msg len
func (this *SMsgHeaderServer2Gate) ServerMsgLen() bsn_common.TMsgLen {
	return this.M_TMsgLenServerMsg
}

func (this *SMsgHeaderServer2Gate) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeaderServe2Gater_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsgHeaderServer2Gate) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	vTMsgLen := this.SMsgHeader.Serialize2Byte(byDatas)

	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], uint16(this.ServerMsgLen()))
	vTMsgLen += 2

	return vTMsgLen
}

func (this *SMsgHeaderServer2Gate) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	vTMsgLen := this.SMsgHeader.DeSerialize(byDatas)

	this.M_TMsgLenServerMsg = bsn_common.TMsgLen(binary.LittleEndian.Uint16(byDatas[vTMsgLen:]))
	vTMsgLen += 2

	return vTMsgLen
}
