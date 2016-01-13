package bsn_log

import (
	"fmt"
	"github.com/bsn069/go/bsn_common"
	"time"
)

func makeLog() ILog {
	pkgName, _, _, _, err := bsn_common.GetCallInfo(2)
	strName := "?"
	if err == nil {
		strName = pkgName
	}

	log := &sLog{
		m_strName:    strName,
		m_u32OutMask: ELevel_All,
		m_u32LogMask: ELevel_All,
		m_timeFunc:   fmtTime,
	}
	return log
}

func fmtTime(t *time.Time) string {
	_, month, day := t.Date()
	hour, minute, second := t.Clock()
	nanosecond := int64(t.Nanosecond()) / (int64)(time.Millisecond)
	return fmt.Sprintf("%02d%02d-%02d:%02d:%02d:%03d", (int)(month), day, hour, minute, second, nanosecond)
}
