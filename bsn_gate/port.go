/*
Package bsn_gate.
*/
package bsn_gate

import (
	"github.com/bsn069/go/bsn_log"
	// "github.com/bsn069/go/bsn_net"
)

type TGroupId uint16

type IGate interface {
}

var GLog = bsn_log.New()

// newGate() IGate
var NewGate = newGate
