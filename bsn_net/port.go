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
type TFuncOnListen func(conn net.Conn) error
type IListen interface {
	SetListenPort(port TPort) error
	SetListenFunc(funcOnListen TFuncOnListen) error
	Listen() (err error)
	StopListen()
}

// (iListenCallBack IListenCallBack) IListen
var NewListen = newListen

type TUserImp interface{}
type TUserId uint32
type IUser interface {
	GetId() TUserId
	GetUserMgr() IUserMgr
	Close()
}
type TID2User map[TUserId]TUserImp

type TFuncNewUser func(user IUser) (TUserImp, error)
type IUserMgr interface {
	SetFuncNewUser(funcNewUser TFuncNewUser) error
	SetListenPort(port TPort) error
	Listen() error
	StopListen()
	GetUser(userId TUserId) TUserImp
	DelUser(userId TUserId)
}

// () (IUserMgr, error)
var NewUserMgr = newUserMgr

var GLog = bsn_log.New()
