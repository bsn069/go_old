package bsn_msg

import (
	"encoding/binary"
	"github.com/bsn069/go/bsn_common"
	"unsafe"
)

type SMsg_Server2Gate_ClientMsg struct {
	M_ClientId uint16
	M_byMsg    []byte
}

func (this *SMsg_Server2Gate_ClientMsg) Fill(vu16ClientId uint16, byMsg []byte) {
	this.M_ClientId = vu16ClientId
	this.M_byMsg = byMsg
}

func (this *SMsg_Server2Gate_ClientMsg) Serialize() []byte {
	var byDatas = make([]byte, int(unsafe.Sizeof(this.M_ClientId))+len(this.M_byMsg))
	this.Serialize2Byte(byDatas)
	return byDatas
}

func (this *SMsg_Server2Gate_ClientMsg) Serialize2Byte(byDatas []byte) bsn_common.TMsgLen {
	var vTMsgLen bsn_common.TMsgLen = 0

	binary.LittleEndian.PutUint16(byDatas[vTMsgLen:], this.M_ClientId)
	vTMsgLen += 2

	copy(byDatas[vTMsgLen:], this.M_byMsg)
	vTMsgLen += bsn_common.TMsgLen(len(this.M_byMsg))

	return vTMsgLen
}

func (this *SMsg_Server2Gate_ClientMsg) DeSerialize(byDatas []byte) bsn_common.TMsgLen {
	var vTMsgLen bsn_common.TMsgLen = 0

	this.M_ClientId = binary.LittleEndian.Uint16(byDatas[vTMsgLen:])
	vTMsgLen += 2

	this.M_byMsg = byDatas[vTMsgLen:]
	vTMsgLen += bsn_common.TMsgLen(len(this.M_byMsg))

	return vTMsgLen
}
