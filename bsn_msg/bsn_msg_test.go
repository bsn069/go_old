package bsn_msg

import (
	// "strconv"
	"testing"
)

func TestBase(t *testing.T) {
	msgHeader := NewMsgHeader(1, 2)
	if msgHeader.Type() != 1 {
		t.Error(msgHeader.Type())
	}
	if msgHeader.Len() != 2 {
		t.Error(msgHeader.Len())
	}

	byDatas := msgHeader.Serialize()
	GSLog.Mustln(byDatas)
	msgHeader.DeSerialize(byDatas)
	if msgHeader.Type() != 1 {
		t.Error(msgHeader.Type())
	}
	if msgHeader.Len() != 2 {
		t.Error(msgHeader.Len())
	}
}
