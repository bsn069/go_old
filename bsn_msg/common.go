package bsn_msg

import (
	"github.com/bsn069/go/bsn_log"
	"unsafe"
)

func init() {
	GSLog = bsn_log.GSLog

	a := new(SMsgHeader)
	if unsafe.Sizeof(*a) != uintptr(CSMsgHeader_Size) {
		panic("CSMsgHeader_Size != sizeof SMsgHeader")
	}
}
