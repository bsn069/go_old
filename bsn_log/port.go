/*
Package bsn_log.
*/
package bsn_log

import (
	"time"
)

type TLevel uint32

const (
	ELevel_Must TLevel = 1 << iota
	ELevel_Debug
	ELevel_Error
	ELevel_Max
	ELevel_All = ELevel_Max - 1
)

type ILog interface {
	SetName(strName string)
	SetOutMask(u32Mask uint32)
	SetLogMask(u32Mask uint32)
	Output(ELevel TLevel, strInfo string)
	Must(v ...interface{})
	Mustln(v ...interface{})
	Mustf(format string, v ...interface{})
	Debug(v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorln(v ...interface{})
	Errorf(format string, v ...interface{})
}

type TTimeFunc func(t *time.Time) string

var New = makeLog
var GLog = New()
