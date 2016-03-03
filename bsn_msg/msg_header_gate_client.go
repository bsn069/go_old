package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	"encoding/binary"
)

const (
	CSMsgHeaderGateClient_Size bsn_common.TMsgLen = CSMsgHeader_Size + 2
)

// msgHeader from client
type SMsgHeaderGateClient struct {
	SMsgHeader
	M_TGateUserId bsn_common.TGateUserId
}

func NewMsgHeaderGateClient(vu16ServerType uint16, vTGateUserId bsn_common.TGateUserId, byMsg []byte) *SMsgHeaderGateClient {
	this := &SMsgHeaderGateClient{}
	this.Fill(vTGateUserId)
	this.SMsgHeader.Fill(bsn_common.TMsgType(vu16ServerType), bsn_common.TMsgLen(len(byMsg)))
	return this
}

func NewMsgHeaderGateClientFromBytes(byDatas []byte) *SMsgHeaderGateClient {
	this := &SMsgHeaderGateClient{}
	this.DeSerialize(byDatas)
	return this
}

func (this *SMsgHeaderGateClient) Fill(vM_TGateUserId bsn_common.TGateUserId) {
	this.M_TGateUserId = vM_TGateUserId
}

func (this *SMsgHeaderGateClient) ServerType() uint16 {
	return uint16(this.Type())
}

func (this *SMsgHeaderGateClient) UserId() bsn_common.TGateUserId {
	return this.M_TGateUserId
}

func (this *SMsgHeaderGateClient) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeaderGateClient_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsgHeaderGateClient) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	vTMsgLen := this.SMsgHeader.Serialize2Byte(byDatas)

	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], uint16(this.UserId()))
	vTMsgLen += 2

	return vTMsgLen
}

func (this *SMsgHeaderGateClient) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	vTMsgLen := this.SMsgHeader.DeSerialize(byDatas)

	this.M_TGateUserId = bsn_common.TGateUserId(binary.LittleEndian.Uint16(byDatas[vTMsgLen:]))
	vTMsgLen += 2

	return vTMsgLen
}
