package bsn_gate

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "runtime"
	// "time"
)

type TServerMsgType uint16

const (
	CServerMsgType_ToUser TServerMsgType = iota + 1
	CServerMsgType_ToGroup
	CServerMsgType_AddUserToGroup
	CServerMsgType_DelUserFromGroup
	CServerMsgType_DelUser

	CSServerMsg_Size bsn_msg.TMsgLen = bsn_msg.CSMsgHeader_Size + 4
)

type SServerMsg struct {
	bsn_msg.SMsgHeader
	M_TGroupId TGroupId
	M_TUserId  TUserId
}

func newServerMsg(vTServerMsgType TServerMsgType, vTGroupId TGroupId, vTUserId TUserId, byMsg []byte) *SServerMsg {
	this := &SServerMsg{}
	this.Fill(vTGroupId, vTUserId)
	this.SMsgHeader.Fill(bsn_msg.TMsgType(vTServerMsgType), bsn_msg.TMsgLen(len(byMsg)))
	return this
}

func (this *SServerMsg) Fill(vTGroupId TGroupId, vTUserId TUserId) {
	this.M_TGroupId = vTGroupId
	this.M_TUserId = vTUserId
}

func (this *SServerMsg) UserId() TUserId {
	return this.M_TUserId
}

func (this *SServerMsg) GroupId() TGroupId {
	return this.M_TGroupId
}

func (this *SServerMsg) Serialize() []byte {
	var byDatas = make([]byte, CSServerMsg_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SServerMsg) Serialize2Byte(byDatas []byte) bsn_msg.TMsgLen {
	vLen := this.SMsgHeader.Serialize2Byte(byDatas)

	bsn_common.WriteUint16(byDatas[vLen:], uint16(this.GroupId()))
	vLen += 2
	bsn_common.WriteUint16(byDatas[vLen:], uint16(this.UserId()))
	vLen += 2

	return vLen
}

func (this *SServerMsg) DeSerialize(byDatas []byte) bsn_msg.TMsgLen {
	vLen := this.SMsgHeader.DeSerialize(byDatas)

	this.M_TGroupId = TGroupId(bsn_common.ReadUint16(byDatas[vLen:]))
	vLen += 2
	this.M_TUserId = TUserId(bsn_common.ReadUint16(byDatas[vLen:]))
	vLen += 2

	return vLen
}
