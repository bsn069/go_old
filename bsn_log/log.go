package bsn_log

import (
	"fmt"
	"time"
)

type sLog struct {
	m_strName string
	// output mask
	m_u32OutMask uint32

	// write to log file mask
	m_u32LogMask uint32
	m_time       time.Time
	m_timeFunc   TTimeFunc
}

func (this *sLog) SetName(strName string) {
	this.m_strName = strName
}

func (this *sLog) SetOutMask(u32Mask uint32) {
	this.m_u32OutMask = u32Mask
}

func (this *sLog) SetLogMask(u32Mask uint32) {
	this.m_u32LogMask = u32Mask
}

func (this *sLog) Output(ELevel uint32, strInfo string) {
	this.m_time = time.Now()
	strTime := this.m_timeFunc(&this.m_time)
	if (this.m_u32OutMask & ELevel) != 0 {
		fmt.Print(strTime + "[" + this.m_strName + "]" + " " + strInfo)
	}
	if (this.m_u32LogMask & ELevel) != 0 {
	}
}

func (this *sLog) Debug(v ...interface{}) {
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Debug, strInfo)
}

func (this *sLog) Debugln(v ...interface{}) {
	v = append(v, "\n")
	this.Debug(v...)
}

func (this *sLog) Debugf(format string, v ...interface{}) {
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Debug, strInfo)
}

func (this *sLog) Error(v ...interface{}) {
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Error, strInfo)
}

func (this *sLog) Errorln(v ...interface{}) {
	v = append(v, "\n")
	this.Error(v...)
}

func (this *sLog) Errorf(format string, v ...interface{}) {
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Error, strInfo)
}

func (this *sLog) Must(v ...interface{}) {
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Must, strInfo)
}

func (this *sLog) Mustln(v ...interface{}) {
	v = append(v, "\n")
	this.Must(v...)
}

func (this *sLog) Mustf(format string, v ...interface{}) {
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Must, strInfo)
}
