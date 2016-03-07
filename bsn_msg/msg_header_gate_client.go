package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	// "encoding/binary"
)

const (
	CSMsgHeaderGateClient_Size bsn_common.TMsgLen = CSMsgHeader_Size + 2
)

// msgHeader from client
type SMsgHeaderGateClient struct {
	SMsgHeader
}

func NewMsgHeaderGateClient(vu8ServerType, vu8ServerId uint8, byMsg []byte) *SMsgHeaderGateClient {
	this := &SMsgHeaderGateClient{}
	vUserId := bsn_common.MakeGateUserId(vu8ServerType, vu8ServerId)
	this.SMsgHeader.Fill(bsn_common.TMsgType(vUserId), bsn_common.TMsgLen(len(byMsg)))
	return this
}

func NewMsgHeaderGateClientFromBytes(byDatas []byte) *SMsgHeaderGateClient {
	this := &SMsgHeaderGateClient{}
	this.DeSerialize(byDatas)
	return this
}

func (this *SMsgHeaderGateClient) UserId() bsn_common.TGateUserId {
	return bsn_common.TGateUserId(this.Type())
}

func (this *SMsgHeaderGateClient) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeaderGateClient_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}
