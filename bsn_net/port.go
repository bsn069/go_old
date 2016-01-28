/*
Package bsn_net.
*/
package bsn_net

import (
	"github.com/bsn069/go/bsn_log"
	"net"
)

type TCloseChan chan bool
type TPort uint16

// (iListenCallBack IListenCallBack) IListen
var NewListen = newListen

type TUserId uint32
type IUser interface {
	GetId() TUserId
	GetUserMgr() IUserMgr
	Close()
	GetConn() net.Conn
}
type TID2User map[TUserId]IUser

// (iUserMgr IUserMgr, userId TUserId, iConn net.Conn) (IUser, error)
var NewUser = newUser

type IUserMgrCallBack interface {
	NewUser(userId TUserId, iConn net.Conn) (IUser, error)
}

type TUserMgrType uint
type IUserMgr interface {
	GetUser(userId TUserId) IUser
	DelUser(userId TUserId)
	Close()
	GetType() TUserMgrType
}

// (userMgrType TUserMgrType, iUserMgrCallBack IUserMgrCallBack) (IUserMgr, error)
var NewUserMgr = newUserMgr

var GLog = bsn_log.New()
