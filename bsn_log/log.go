package bsn_log

import (
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"runtime"
	"time"
)

var fl_u32Id uint32 = 0

type sLog struct {
	m_strName string
	// output mask
	m_u32OutMask uint32

	// write to log file mask
	m_u32LogMask   uint32
	m_time         time.Time
	m_timeFmtFunc  TTimeFmtFunc
	m_outFmtFunc   TOutFmtFunc
	m_debugFmtFunc TDebugFmtFunc
	m_depth        int
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

func (this *sLog) SetTimeFmtFunc(timeFmtFunc TTimeFmtFunc) {
	this.m_timeFmtFunc = timeFmtFunc
}

func (this *sLog) SetOutFmtFunc(outFmtFunc TOutFmtFunc) {
	this.m_outFmtFunc = outFmtFunc
}

func (this *sLog) SetDebugFmtFunc(debugFmtFunc TDebugFmtFunc) {
	this.m_debugFmtFunc = debugFmtFunc
}

func (this *sLog) FuncGuard() {
	if err := recover(); err != nil {
		buf := make([]byte, 2048)
		len := runtime.Stack(buf, false)
		this.Errorln(string(buf[:len]))
	}
}

func (this *sLog) Output(level TLevel, strInfo string) {
	defer bsn_common.FuncGuard()
	fl_u32Id++
	this.m_depth++
	this.m_time = time.Now()
	strTime := this.m_timeFmtFunc(&this.m_time)
	strDebugInfo := this.m_debugFmtFunc(this.m_depth)
	if (this.m_u32OutMask & uint32(level)) != 0 {
		strOutInfo := this.m_outFmtFunc(level, &strTime, &this.m_strName, &strInfo, &strDebugInfo, fl_u32Id)
		fmt.Print(strOutInfo)
	}
	if (this.m_u32LogMask & uint32(level)) != 0 {

	}
	this.m_depth = 0
}

func (this *sLog) Debug(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Debug, strInfo)
}

func (this *sLog) Debugln(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	v = append(v, "\n")
	this.Debug(v...)
}

func (this *sLog) Debugf(format string, v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Debug, strInfo)
}

func (this *sLog) Error(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Error, strInfo)
}

func (this *sLog) Errorln(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	v = append(v, "\n")
	this.Error(v...)
}

func (this *sLog) Errorf(format string, v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Error, strInfo)
}

func (this *sLog) Must(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Must, strInfo)
}

func (this *sLog) Mustln(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	v = append(v, "\n")
	this.Must(v...)
}

func (this *sLog) Mustf(format string, v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.m_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Must, strInfo)
}
