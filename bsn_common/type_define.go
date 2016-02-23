package bsn_common

import (
	"net"
	"time"
)

type TVoid interface{}

type TMsgType uint16
type TMsgLen uint16

type TNetChanClose chan bool
type TNetChanConn chan net.Conn

type TInputParams []string
type TInputUpperName2CmdStruct map[string]TVoid
type TInputUpperName2RegName map[string]string

// [modName][cmdUpper] = help
type TInputHelp map[string]map[string]string

type TGateId2Gate map[TGateId]TVoid
type TGateServerMsgType uint16
type TGateGroupId uint16
type TGateId uint32
type TGateUserId uint16
type TGateUserId2User map[TGateUserId]IGateUser
type TGateUserMgrType uint32
type IGateUser interface {
	GetId() TGateUserId
	GetConn() net.Conn
	ReadMsgHeader() error
	ReadMsgBody() error
}
type IGateUserMgr interface {
	AddUser(vIUser IGateUser)
	DelUser(vIUser IGateUser)
	GetUser(vTUserId TGateUserId) IGateUser
	GetType() TGateUserMgrType
}

type TLogLevel uint32
type TLogTimeFmtFunc func(t *time.Time) string
type TLogDebugFmtFunc func(depth int) string
type TLogOutFmtFunc func(level TLogLevel, strTime, strInfo, strDebugInfo *string, id uint32) string

const (
	ELogLevel_Must TLogLevel = 1 << iota
	ELogLevel_Debug
	ELogLevel_Error
	ELogLevel_Max
	ELogLevel_All = ELogLevel_Max - 1
)

// String returns the English name of the level ("Debug", "Must ", ...).
func (level TLogLevel) String() string {
	switch level {
	case ELogLevel_Debug:
		return "Debug"
	case ELogLevel_Must:
		return "Must "
	case ELogLevel_Error:
		return "Error"
	default:
		return "     "
	}
}
