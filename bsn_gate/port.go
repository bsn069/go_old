/*
Package bsn_gate.
*/
package bsn_gate

import (
	"github.com/bsn069/go/bsn_log"
)

type IGate interface {
	GetClientMgr() IClientMgr
}

var GLog = bsn_log.New()

// newGate() IGate
var NewGate = newGate
