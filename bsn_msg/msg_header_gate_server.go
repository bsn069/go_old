package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	"encoding/binary"
)

const (
	CServerMsgType_ToUser bsn_common.TGateServerMsgType = iota + 1
	CServerMsgType_ToGroup
	CServerMsgType_AddUserToGroup
	CServerMsgType_DelUserFromGroup
	CServerMsgType_DelUser

	CSMsgHeaderGateServer_Size bsn_common.TMsgLen = CSMsgHeader_Size + 4
)

// msgHeader from server
type SMsgHeaderGateServer struct {
	SMsgHeader
	M_TGateGroupId bsn_common.TGateGroupId
	M_TGateUserId  bsn_common.TGateUserId
}

func NewMsgHeaderGateServer(vTGateServerMsgType bsn_common.TGateServerMsgType, vTGroupId bsn_common.TGateGroupId, vTGateUserId bsn_common.TGateUserId, byMsg []byte) *SMsgHeaderGateServer {
	this := &SMsgHeaderGateServer{}
	this.Fill(vTGroupId, vTGateUserId)
	this.SMsgHeader.Fill(bsn_common.TMsgType(vTGateServerMsgType), bsn_common.TMsgLen(len(byMsg)))
	return this
}

func NewMsgHeaderGateServerFromBytes(byDatas []byte) *SMsgHeaderGateServer {
	this := &SMsgHeaderGateServer{}
	this.DeSerialize(byDatas)
	return this
}

func (this *SMsgHeaderGateServer) Fill(vTGroupId bsn_common.TGateGroupId, vTGateUserId bsn_common.TGateUserId) {
	this.M_TGateGroupId = vTGroupId
	this.M_TGateUserId = vTGateUserId
}

func (this *SMsgHeaderGateServer) ServerMsgType() bsn_common.TGateServerMsgType {
	return bsn_common.TGateServerMsgType(this.Type())
}

func (this *SMsgHeaderGateServer) UserId() bsn_common.TGateUserId {
	return this.M_TGateUserId
}

func (this *SMsgHeaderGateServer) GroupId() bsn_common.TGateGroupId {
	return this.M_TGateGroupId
}

func (this *SMsgHeaderGateServer) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeaderGateServer_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsgHeaderGateServer) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	vTMsgLen := this.SMsgHeader.Serialize2Byte(byDatas)

	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], uint16(this.GroupId()))
	vTMsgLen += 2
	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], uint16(this.UserId()))
	vTMsgLen += 2

	return vTMsgLen
}

func (this *SMsgHeaderGateServer) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	vTMsgLen := this.SMsgHeader.DeSerialize(byDatas)

	this.M_TGateGroupId = bsn_common.TGateGroupId(binary.LittleEndian.Uint16(byDatas[vTMsgLen:]))
	vTMsgLen += 2
	this.M_TGateUserId = bsn_common.TGateUserId(binary.LittleEndian.Uint16(byDatas[vTMsgLen:]))
	vTMsgLen += 2

	return vTMsgLen
}
