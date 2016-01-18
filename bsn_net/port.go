/*
Package bsn_net.
*/
package bsn_net

import (
	"github.com/bsn069/go/bsn_log"
	"net"
)

type IListen interface {
	SetListenPort(u16Port uint16) error
	SetListenFunc(funcOnListen TFuncOnListen) error
	Listen() (err error)
	StopListen()
}

var GLog = bsn_log.New()

// () *sListen
var NewListen = newListen

type TFuncOnListen func(conn net.Conn)
