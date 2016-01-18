/*
Package bsn_gate.
*/
package bsn_gate

import (
	"github.com/bsn069/go/bsn_log"
)

type IClientMgr interface {
	SetListenPort(u16Port uint16) error
	Listen() error
	StopListen()
}

type IGate interface {
	GetClientMgr() IClientMgr
}

var GLog = bsn_log.New()

// newGate() IGate
var NewGate = newGate
