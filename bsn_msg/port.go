/*
Package bsn_msg.
*/
package bsn_msg

import (
	"github.com/bsn069/go/bsn_log"
)

type IMsgHeader interface {
	Len() uint16
	Type() uint16
	DeSerialize(data []byte)
	Serialize() []byte
	Serialize2Byte(byDatas []byte)
}

var GLog = bsn_log.New()

// (u16Type, u16Len uint16) IMsgHeader
var NewMsgHeader = newMsgHeader
