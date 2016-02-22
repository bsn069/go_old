package bsn_gate

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	// "runtime"
	// "time"
)

const (
	CServerMsgType_ToUser bsn_common.TGateServerMsgType = iota + 1
	CServerMsgType_ToGroup
	CServerMsgType_AddUserToGroup
	CServerMsgType_DelUserFromGroup
	CServerMsgType_DelUser

	CSServerMsg_Size bsn_common.TMsgLen = bsn_msg.CSMsgHeader_Size + 4
)

type SServerMsg struct {
	bsn_msg.SMsgHeader
	M_TGroupId bsn_common.TGateGroupId
	M_TUserId  bsn_common.TGateUserId
}

func newServerMsg(vTGateServerMsgType bsn_common.TGateServerMsgType, vTGroupId bsn_common.TGateGroupId, vTUserId bsn_common.TGateUserId, byMsg []byte) *SServerMsg {
	this := &SServerMsg{}
	this.Fill(vTGroupId, vTUserId)
	this.SMsgHeader.Fill(bsn_common.TMsgType(vTGateServerMsgType), bsn_common.TMsgLen(len(byMsg)))
	return this
}

func (this *SServerMsg) Fill(vTGroupId bsn_common.TGateGroupId, vTUserId bsn_common.TGateUserId) {
	this.M_TGroupId = vTGroupId
	this.M_TUserId = vTUserId
}

func (this *SServerMsg) UserId() bsn_common.TGateUserId {
	return this.M_TUserId
}

func (this *SServerMsg) GroupId() bsn_common.TGateGroupId {
	return this.M_TGroupId
}

func (this *SServerMsg) Serialize() []byte {
	var byDatas = make([]byte, CSServerMsg_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SServerMsg) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	vLen := this.SMsgHeader.Serialize2Byte(byDatas)

	bsn_common.WriteUint16(byDatas[vLen:], uint16(this.GroupId()))
	vLen += 2
	bsn_common.WriteUint16(byDatas[vLen:], uint16(this.UserId()))
	vLen += 2

	return vLen
}

func (this *SServerMsg) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	vLen := this.SMsgHeader.DeSerialize(byDatas)

	this.M_TGroupId = bsn_common.TGateGroupId(bsn_common.ReadUint16(byDatas[vLen:]))
	vLen += 2
	this.M_TUserId = bsn_common.TGateUserId(bsn_common.ReadUint16(byDatas[vLen:]))
	vLen += 2

	return vLen
}
