package bsn_msg

import (
// "fmt"
// "github.com/bsn069/go/bsn_common"
// "runtime"
// "time"
)

type SMsg struct {
	M_byMsg      []byte
	M_SMsgHeader SMsgHeader
}

func NewSMsg() (this *SMsg) {
	this = &SMsg{}
	return
}

func (this *SMsg) MsgBodyBuffer(byMsgHeader []byte) (byMsgBody []byte) {
	this.M_SMsgHeader.DeSerialize(byMsgHeader)
	vLen := this.M_SMsgHeader.Len()
	this.M_byMsg = make([]byte, vLen+CSMsgHeader_Size)
	copy(this.M_byMsg, byMsgHeader)
	byMsgBody = this.M_byMsg[CSMsgHeader_Size:]
	return
}
