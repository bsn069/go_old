/*
Package bsn_msg.
*/
package bsn_msg

import (
	"github.com/bsn069/go/bsn_log"
)

type TMsgType uint16
type TMsgLen uint16

const (
	CSMsgHeader_Size TMsgLen = 4
)

var GLog = bsn_log.New()

// (vTMsgType TMsgType, vTMsgLen TMsgLen) *SMsgHeader
var NewMsgHeader = newMsgHeader
