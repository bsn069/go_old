package bsn_msg

import (
	// "fmt"
	// "github.com/bsn069/go/bsn_common"
	// "runtime"
	// "time"
	"sync"
)

var gSMsgPool sync.Pool

type SMsg struct {
	SMsgHeader
	M_byMsg []byte
}

func NewSMsg() (this *SMsg) {
	poolObj := gSMsgPool.Get()
	this = poolObj.(*SMsg)
	return this
}

func (this *SMsg) Del() {
	gSMsgPool.Put(this)
}

func (this *SMsg) Init(byMsgHeader []byte) {
	this.DeSerialize(byMsgHeader)

	vLen := this.Len() + CSMsgHeader_Size
	if int(vLen) > cap(this.M_byMsg) {
		this.M_byMsg = make([]byte, vLen)
	} else {
		this.M_byMsg = this.M_byMsg[0:vLen]
	}

	copy(this.M_byMsg, byMsgHeader)
}

func (this *SMsg) MsgBodyBuffer() (byMsgBody []byte) {
	byMsgBody = this.M_byMsg[CSMsgHeader_Size:]
	return byMsgBody
}
