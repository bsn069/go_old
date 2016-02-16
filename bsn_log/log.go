package bsn_log

import (
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"runtime"
	"time"
)

type TLevel uint32
type TTimeFmtFunc func(t *time.Time) string
type TDebugFmtFunc func(depth int) string
type TOutFmtFunc func(level TLevel, strTime, strModName, strInfo, strDebugInfo *string, id uint32) string

const (
	ELevel_Must TLevel = 1 << iota
	ELevel_Debug
	ELevel_Error
	ELevel_Max
	ELevel_All = ELevel_Max - 1
)

var fl_u32Id uint32 = 0

type SLog struct {
	M_strName string
	// output mask
	M_u32OutMask uint32

	// write to log file mask
	M_u32LogMask   uint32
	M_time         time.Time
	M_timeFmtFunc  TTimeFmtFunc
	M_outFmtFunc   TOutFmtFunc
	M_debugFmtFunc TDebugFmtFunc
	M_depth        int
	M_SCmd         SCmd
}

func (this *SLog) SetName(strName string) {
	this.M_strName = strName
}

func (this *SLog) SetOutMask(u32Mask uint32) {
	this.M_u32OutMask = u32Mask
}

func (this *SLog) SetLogMask(u32Mask uint32) {
	this.M_u32LogMask = u32Mask
}

func (this *SLog) SetTimeFmtFunc(timeFmtFunc TTimeFmtFunc) {
	this.M_timeFmtFunc = timeFmtFunc
}

func (this *SLog) SetOutFmtFunc(outFmtFunc TOutFmtFunc) {
	this.M_outFmtFunc = outFmtFunc
}

func (this *SLog) SetDebugFmtFunc(debugFmtFunc TDebugFmtFunc) {
	this.M_debugFmtFunc = debugFmtFunc
}

func (this *SLog) FuncGuard() {
	if err := recover(); err != nil {
		buf := make([]byte, 2048)
		len := runtime.Stack(buf, false)
		this.Errorln(string(buf[:len]))
	}
}

func (this *SLog) Output(level TLevel, strInfo string) {
	defer bsn_common.FuncGuard()
	fl_u32Id++
	this.M_depth++
	this.M_time = time.Now()
	strTime := this.M_timeFmtFunc(&this.M_time)
	strDebugInfo := this.M_debugFmtFunc(this.M_depth)
	if (this.M_u32OutMask & uint32(level)) != 0 {
		strOutInfo := this.M_outFmtFunc(level, &strTime, &this.M_strName, &strInfo, &strDebugInfo, fl_u32Id)
		fmt.Print(strOutInfo)
	}
	if (this.M_u32LogMask & uint32(level)) != 0 {

	}
	this.M_depth = 0
}

func (this *SLog) Debug(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Debug, strInfo)
}

func (this *SLog) Debugln(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	v = append(v, "\n")
	this.Debug(v...)
}

func (this *SLog) Debugf(format string, v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Debug, strInfo)
}

func (this *SLog) Error(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Error, strInfo)
}

func (this *SLog) Errorln(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	v = append(v, "\n")
	this.Error(v...)
}

func (this *SLog) Errorf(format string, v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Error, strInfo)
}

func (this *SLog) Must(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprint(v...)
	this.Output(ELevel_Must, strInfo)
}

func (this *SLog) Mustln(v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	v = append(v, "\n")
	this.Must(v...)
}

func (this *SLog) Mustf(format string, v ...interface{}) {
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.Output(ELevel_Must, strInfo)
}
