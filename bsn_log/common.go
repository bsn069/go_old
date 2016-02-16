package bsn_log

import (
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"path"
	// "strings"
	"time"
)

// String returns the English name of the level ("Debug", "Must ", ...).
func (level TLevel) String() string {
	switch level {
	case ELevel_Debug:
		return "Debug"
	case ELevel_Must:
		return "Must "
	case ELevel_Error:
		return "Error"
	default:
		return "     "
	}
}

func New() *SLog {
	pkgName, _, _, _, err := bsn_common.GetCallInfo(2)
	strName := "?"
	if err == nil {
		strName = pkgName
	}

	log := &SLog{
		M_strName:      strName,
		M_u32OutMask:   uint32(ELevel_All),
		M_u32LogMask:   uint32(ELevel_All),
		M_timeFmtFunc:  fmtTime,
		M_outFmtFunc:   fmtOut,
		M_debugFmtFunc: fmtDebug,
	}
	log.M_SCmd.M_SLog = log

	return log
}

func fmtTime(t *time.Time) string {
	// _, month, day := t.Date()
	hour, minute, second := t.Clock()
	// nanosecond := int64(t.Nanosecond()) / (int64)(time.Millisecond)
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func fmtOut(level TLevel, strTime, strModName, strInfo, strDebugInfo *string, id uint32) string {
	return fmt.Sprintf("[%v][%v][%v][%v][%v]%v", level, *strTime, id, *strModName, *strDebugInfo, *strInfo)
}

func fmtDebug(depth int) string {
	file, line := bsn_common.GetCallerFileLine(depth + 2)
	file = path.Base(file)
	return fmt.Sprintf("%v:%v", file, line)
}
