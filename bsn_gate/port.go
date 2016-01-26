/*
Package bsn_gate.
*/
package bsn_gate

import (
	"github.com/bsn069/go/bsn_log"
	"github.com/bsn069/go/bsn_net"
)

type IGate interface {
	GetServerMgr() IServerMgr
}

var GLog = bsn_log.New()

// newGate() IGate
var NewGate = newGate

type IServerMgr interface {
	bsn_net.IUserMgr
	bsn_net.IUserMgrCallBack
}
