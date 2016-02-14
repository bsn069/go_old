package bsn_gate

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "runtime"
	// "time"
)

type TClientMsgType uint16

const (
	CClientMsgType_ToUser TClientMsgType = iota + 1

	CSClientMsg_Size bsn_msg.TMsgLen = bsn_msg.CSMsgHeader_Size + 2
)

type SClientMsg struct {
	bsn_msg.SMsgHeader
	M_TUserId TUserId
}

func newClientMsg(vTClientMsgType TClientMsgType, vTGroupId TGroupId, vTUserId TUserId, byMsg []byte) *SClientMsg {
	this := &SClientMsg{}
	this.Fill(vTUserId)
	this.SMsgHeader.Fill(bsn_msg.TMsgType(vTClientMsgType), bsn_msg.TMsgLen(len(byMsg)))
	return this
}

func (this *SClientMsg) Fill(vTUserId TUserId) {
	this.M_TUserId = vTUserId
}

func (this *SClientMsg) UserId() TUserId {
	return this.M_TUserId
}

func (this *SClientMsg) Serialize() []byte {
	var byDatas = make([]byte, CSClientMsg_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SClientMsg) Serialize2Byte(byDatas []byte) bsn_msg.TMsgLen {
	vLen := this.SMsgHeader.Serialize2Byte(byDatas)

	bsn_common.WriteUint16(byDatas[vLen:], uint16(this.UserId()))
	vLen += 2

	return vLen
}

func (this *SClientMsg) DeSerialize(byDatas []byte) bsn_msg.TMsgLen {
	vLen := this.SMsgHeader.DeSerialize(byDatas)

	this.M_TUserId = TUserId(bsn_common.ReadUint16(byDatas[vLen:]))
	vLen += 2

	return vLen
}
