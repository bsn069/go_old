/*
Package bsn_net.
*/
package bsn_net

import (
	"github.com/bsn069/go/bsn_log"
	"net"
)

type IListenCallBack interface {
	OnListen(conn net.Conn) error
}

type TCloseChan chan bool
type TPort uint16
type IListen interface {
	SetListenPort(port TPort) error
	Listen() (err error)
	StopListen()
}

// (iListenCallBack IListenCallBack) IListen
var NewListen = newListen

type TUserId uint32
type IUser interface {
	GetId() TUserId
	GetUserMgr() IUserMgr
	Close()
}
type TID2User map[TUserId]IUser

// (userId TUserId, iConn net.Conn) (IUser, error)
var NewUser = newUser

type IUserMgrCallBack interface {
	NewUser(userId TUserId, iConn net.Conn) (IUser, error)
}

type IUserMgr interface {
	IListen
	GetUser(userId TUserId) IUser
	DelUser(userId TUserId)
}

// (iUserMgrCallBack IUserMgrCallBack) (IUserMgr, error)
var NewUserMgr = newUserMgr

var GLog = bsn_log.New()
