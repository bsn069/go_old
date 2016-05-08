package bsn_log

import (
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"runtime"
	"sync"
	"time"
)

var fl_u32Id uint32 = 0

type SLog struct {
	// output mask
	M_u32OutMask uint32

	// write to log file mask
	M_u32LogMask   uint32
	M_time         time.Time
	M_timeFmtFunc  bsn_common.TLogTimeFmtFunc
	M_outFmtFunc   bsn_common.TLogOutFmtFunc
	M_debugFmtFunc bsn_common.TLogDebugFmtFunc
	M_depth        int

	M_Mutex sync.Mutex
}

func (this *SLog) SetOutMask(u32Mask uint32) {
	this.M_u32OutMask = u32Mask
}

func (this *SLog) SetLogMask(u32Mask uint32) {
	this.M_u32LogMask = u32Mask
}

func (this *SLog) SetTimeFmtFunc(timeFmtFunc bsn_common.TLogTimeFmtFunc) {
	this.M_timeFmtFunc = timeFmtFunc
}

func (this *SLog) SetOutFmtFunc(outFmtFunc bsn_common.TLogOutFmtFunc) {
	this.M_outFmtFunc = outFmtFunc
}

func (this *SLog) SetDebugFmtFunc(debugFmtFunc bsn_common.TLogDebugFmtFunc) {
	this.M_debugFmtFunc = debugFmtFunc
}

func (this *SLog) FuncGuard() {
	if err := recover(); err != nil {
		buf := make([]byte, 2048)
		len := runtime.Stack(buf, false)
		this.Errorln(string(buf[:len]))
	}
}

func (this *SLog) output(level bsn_common.TLogLevel, strInfo string) {
	defer bsn_common.FuncGuard()
	fl_u32Id++
	this.M_depth++
	this.M_time = time.Now()
	strTime := this.M_timeFmtFunc(&this.M_time)
	strDebugInfo := this.M_debugFmtFunc(this.M_depth)
	if (this.M_u32OutMask & uint32(level)) != 0 {
		strOutInfo := this.M_outFmtFunc(level, &strTime, &strInfo, &strDebugInfo, fl_u32Id)
		fmt.Print(strOutInfo)
	}
	if (this.M_u32LogMask & uint32(level)) != 0 {

	}
	this.M_depth = 0
}

func (this *SLog) Debug(v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprint(v...)
	this.output(bsn_common.ELogLevel_Debug, strInfo)
}

func (this *SLog) Debugln(v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	v = append(v, "\n")
	strInfo := fmt.Sprint(v...)
	this.output(bsn_common.ELogLevel_Debug, strInfo)
}

func (this *SLog) Debugf(format string, v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.output(bsn_common.ELogLevel_Debug, strInfo)
}

func (this *SLog) Error(v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprint(v...)
	this.output(bsn_common.ELogLevel_Error, strInfo)
}

func (this *SLog) Errorln(v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	v = append(v, "\n")
	strInfo := fmt.Sprint(v...)
	this.output(bsn_common.ELogLevel_Error, strInfo)
}

func (this *SLog) Errorf(format string, v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.output(bsn_common.ELogLevel_Error, strInfo)
}

func (this *SLog) Must(v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprint(v...)
	this.output(bsn_common.ELogLevel_Must, strInfo)
}

func (this *SLog) Mustln(v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	v = append(v, "\n")
	strInfo := fmt.Sprint(v...)
	this.output(bsn_common.ELogLevel_Must, strInfo)
}

func (this *SLog) Mustf(format string, v ...interface{}) {
	defer this.M_Mutex.Unlock()
	this.M_Mutex.Lock()
	defer bsn_common.FuncGuard()
	this.M_depth++
	strInfo := fmt.Sprintf(format, v...)
	this.output(bsn_common.ELogLevel_Must, strInfo)
}
