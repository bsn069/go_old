/*
Package bsn_gate.
*/
package bsn_gate

import (
	"github.com/bsn069/go/bsn_log"
	"github.com/bsn069/go/bsn_net"
)

const (
	CClientMgr bsn_net.TUserMgrType = iota + 1
	CServerMgr
)

type IUser interface {
	bsn_net.IUser
}

type IUserMgr interface {
	bsn_net.IUserMgr
	bsn_net.IUserMgrCallBack
}

type IGate interface {
	GetServerMgr() IUserMgr
	GetClientMgr() IUserMgr
	Listen()
	StopListen()
	Close()
}

var GLog = bsn_log.New()

// newGate() IGate
var NewGate = newGate
