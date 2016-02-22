package bsn_log

import (
	// "strconv"
	"github.com/bsn069/go/bsn_common"
	"testing"
)

func TestBase(t *testing.T) {
	GLog, _ := New()

	GLog.SetOutMask(uint32(bsn_common.ELogLevel_All))
	GLog.Must("this is Must")
	GLog.Error("this is Error")
	GLog.Debug("this is Debug")
	GLog.Output(bsn_common.ELogLevel_All, "==\n")

	GLog.SetOutMask(uint32(bsn_common.ELogLevel_Must))
	GLog.Mustln("this is Must")
	GLog.Errorln("this is Error")
	GLog.Debugln("this is Debug")
	GLog.Output(bsn_common.ELogLevel_All, "==\n")

	GLog.SetOutMask(uint32(bsn_common.ELogLevel_Error))
	GLog.Mustf("this is Must")
	GLog.Errorf("this is Error")
	GLog.Debugf("this is Debug")
	GLog.Output(bsn_common.ELogLevel_All, "==\n")

	GLog.SetOutMask(uint32(bsn_common.ELogLevel_Debug))
	GLog.Mustf("this is Must %v \n", 1)
	GLog.Errorf("this is Error %v \n", 2)
	GLog.Debugf("this is Debug %v \n", 3)
	GLog.Output(bsn_common.ELogLevel_All, "==\n")
}
