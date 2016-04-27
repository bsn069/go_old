package bsn_msg

import (
	"encoding/binary"
	"github.com/bsn069/go/bsn_common"
	"unsafe"
)

type SMsg_Gate2Server_ClientMsg struct {
	M_ClientId   uint16
	M_SMsgHeader SMsgHeader
	M_byMsgBody  []byte
}

func (this *SMsg_Gate2Server_ClientMsg) Fill(vu16ClientId uint16, vSMsgHeader *SMsgHeader, byMsgBody []byte) {
	this.M_ClientId = vu16ClientId
	this.M_SMsgHeader = *vSMsgHeader
	this.M_byMsgBody = byMsgBody
}

func (this *SMsg_Gate2Server_ClientMsg) Serialize() []byte {
	var byDatas = make([]byte, int(unsafe.Sizeof(this.M_ClientId))+int(unsafe.Sizeof(this.M_SMsgHeader))+len(this.M_byMsgBody))
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsg_Gate2Server_ClientMsg) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	var vTMsgLen bsn_common.TMsgLen = 0

	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], this.M_ClientId)
	vTMsgLen += 2

	vTMsgLen += this.M_SMsgHeader.Serialize2Byte(byDatas[vTMsgLen:])

	copy(byDatas[vTMsgLen:], this.M_byMsgBody)
	vTMsgLen += bsn_common.TMsgLen(len(this.M_byMsgBody))

	return vTMsgLen
}

func (this *SMsg_Gate2Server_ClientMsg) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	var vTMsgLen bsn_common.TMsgLen = 0

	this.M_ClientId = binary.LittleEndian.Uint16(byDatas[vTMsgLen:])
	vTMsgLen += 2

	vTMsgLen += this.M_SMsgHeader.DeSerialize(byDatas[vTMsgLen:])

	this.M_byMsgBody = byDatas[vTMsgLen:]
	vTMsgLen += bsn_common.TMsgLen(len(this.M_byMsgBody))

	return vTMsgLen
}
