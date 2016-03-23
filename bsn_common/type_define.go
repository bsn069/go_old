package bsn_common

import (
	"net"
	"time"
)

type TVoid interface{}

type TMsgType uint16
type TMsgLen uint16

type IMsgSender interface {
	Send(byData []byte) error
}

type TNetChanClose chan bool
type TNetChanConn chan net.Conn

type TInputParams []string
type TInputUpperName2CmdStruct map[string]TVoid
type TInputUpperName2RegName map[string]string

// [modName][cmdUpper] = help
type TInputHelp map[string]map[string]string
type TInputOK chan bool

type TGateId2Gate map[TGateId]TVoid
type TGateServerMsgType uint16
type TGateGroupId uint16
type TGateId uint32
type TGateUserId uint16
type TGateUserId2User map[TGateUserId]IGateUser
type TGateUserMgrType uint32
type IGateUser interface {
	Id() TGateUserId
	Conn() net.Conn
	ReadMsgHeader() error
	ReadMsgBody() error
	Send(byData []byte) error
	Close() error
}
type IGateUserMgr interface {
	AddUser(vIUser IGateUser) error
	DelUser(vIUser IGateUser) error
	User(vTUserId TGateUserId) (IGateUser, error)
	Type() TGateUserMgrType
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

func MakeGateUserId(vu8ServerType, vu8ServerId uint8) TGateUserId {
	var vu16UserId uint16
	vu16UserId = uint16(vu8ServerType)
	vu16UserId <<= 8
	vu16UserId += uint16(vu8ServerId)
	return TGateUserId(vu16UserId)
}
