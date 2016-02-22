package bsn_log

import (
	"fmt"
	"github.com/bsn069/go/bsn_common"
	// "path"
	// "strings"
	"time"
)

func init() {
	GSLog, GSCmd = New()
}

func New() (*SLog, *SCmd) {
	vSLog := &SLog{
		M_u32OutMask:   uint32(bsn_common.ELogLevel_All),
		M_u32LogMask:   uint32(bsn_common.ELogLevel_All),
		M_timeFmtFunc:  fmtTime,
		M_outFmtFunc:   fmtOut,
		M_debugFmtFunc: fmtDebug,
	}

	return vSLog, &SCmd{M_SLog: vSLog}
}

func fmtTime(t *time.Time) string {
	// _, month, day := t.Date()
	hour, minute, second := t.Clock()
	// nanosecond := int64(t.Nanosecond()) / (int64)(time.Millisecond)
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func fmtOut(level bsn_common.TLogLevel, strTime, strInfo, strDebugInfo *string, id uint32) string {
	return fmt.Sprintf("[%v][%v][%v][%v]%v", level, *strTime, id, *strDebugInfo, *strInfo)
}

func fmtDebug(depth int) string {
	file, line := bsn_common.GetCallerFileLine(depth + 2)
	pkgName, fileName, _ := bsn_common.GetPkgFileName(file)
	return fmt.Sprintf("%v/%v:%v", pkgName, fileName, line)
}
