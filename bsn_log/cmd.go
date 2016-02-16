package bsn_log

import (
// "github.com/bsn069/go/bsn_common"
)

type SCmd struct {
	M_SLog *SLog
}

func (this *SCmd) TEST(strArray []string) {

}

func (this *SCmd) TEST_help(vTParams []string) {
	GLog.Mustln("    TEST_help")
}
