package bsn_msg

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	// "encoding/binary"
	"net"
	"unsafe"
)

const (
	CSMsgHeaderGateClient_Size bsn_common.TMsgLen = CSMsgHeader_Size
)

// msgHeader from client
type SMsgHeaderGateClient struct {
	SMsgHeader
}

func NewMsgHeaderGateClient(vu8ServerType, vu8ServerId uint8, vTMsgLen bsn_common.TMsgLen) *SMsgHeaderGateClient {
	this := &SMsgHeaderGateClient{}
	vUserId := bsn_common.MakeGateUserId(vu8ServerType, vu8ServerId)
	this.SMsgHeader.Fill(bsn_common.TMsgType(vUserId), vTMsgLen)
	return this
}

func NewMsgHeaderGateClientFromBytes(byDatas []byte) *SMsgHeaderGateClient {
	this := &SMsgHeaderGateClient{}
	this.DeSerialize(byDatas)
	return this
}

func NewClient2GateMsg(vu8ServerType, vu8ServerId uint8, u16Type uint16, byMsg []byte) []byte {
	vTMsgLen := bsn_common.TMsgLen(len(byMsg))
	vSMsgHeader := NewMsgHeader(bsn_common.TMsgType(u16Type), vTMsgLen)
	vSMsgHeaderGateClient := NewMsgHeaderGateClient(vu8ServerType, vu8ServerId, vTMsgLen+bsn_common.TMsgLen(unsafe.Sizeof(*vSMsgHeader)))

	byData := make([]byte, vSMsgHeaderGateClient.Len()+bsn_common.TMsgLen(unsafe.Sizeof(*vSMsgHeaderGateClient)))
	copy(byData, vSMsgHeaderGateClient.Serialize())
	copy(byData[int(unsafe.Sizeof(*vSMsgHeaderGateClient)):], vSMsgHeader.Serialize())
	copy(byData[int(unsafe.Sizeof(*vSMsgHeaderGateClient))+int(unsafe.Sizeof(*vSMsgHeader)):], byMsg)
	GSLog.Debugln("byData= ", byData)
	return byData
}

func SendClient2GateMsg(conn net.Conn, vu8ServerType, vu8ServerId uint8, u16Type uint16, byMsg []byte) (err error) {
	byData := NewClient2GateMsg(vu8ServerType, vu8ServerId, u16Type, byMsg)
	_, err = conn.Write(byData)
	if err != nil {
		return
	}

	return
}

func (this *SMsgHeaderGateClient) UserId() bsn_common.TGateUserId {
	return bsn_common.TGateUserId(this.Type())
}

func (this *SMsgHeaderGateClient) Serialize() []byte {
	var byDatas = make([]byte, CSMsgHeaderGateClient_Size)
	this.Serialize2Byte(byDatas)
	return byDatas
}
