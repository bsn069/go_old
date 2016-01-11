/*
Package bsn_log.
*/
package bsn_log

import (
	// "fmt"
	"github.com/bsn069/go/bsn_common"
	"path"
)

const (
	ELevel_Must uint32 = 1 << iota
	ELevel_Debug
	ELevel_Error
	ELevel_Max
	ELevel_All = ELevel_Max - 1
)

type ILog interface {
	SetName(strName string)
	SetOutMask(u32Mask uint32)
	SetLogMask(u32Mask uint32)
	Output(ELevel uint32, strInfo string)
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

func New(strName string) ILog {
	if strName == "" {
		filePath, _ := bsn_common.GetCallerFileLine(2)
		dir := path.Dir(filePath)
		pkg := path.Base(dir)
		strName = pkg
	}

	log := &sLog{
		m_strName:    strName,
		m_u32OutMask: ELevel_All,
		m_u32LogMask: ELevel_All,
	}
	return log
}
