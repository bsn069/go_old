package bsn_log

import (
	"github.com/bsn069/go/bsn_common"
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
	}
	return log
}
