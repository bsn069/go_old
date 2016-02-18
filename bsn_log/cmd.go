package bsn_log

import (
	"github.com/bsn069/go/bsn_common"
	"strconv"
)

type SCmd struct {
	M_SLog *SLog
}

func (this *SCmd) TEST(vTInputParams bsn_common.TInputParams) {

}

func (this *SCmd) TEST_help(vTInputParams bsn_common.TInputParams) {
	this.M_SLog.Mustln("    TEST_help")
}

func (this *SCmd) SETOUTMASK(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		this.M_SLog.Errorln("param must a number")
		return
	}

	uMask, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		this.M_SLog.Errorln(err)
		return
	}

	this.M_SLog.SetOutMask(uint32(uMask))
}

func (this *SCmd) SETOUTMASK_help(vTInputParams bsn_common.TInputParams) {
	this.M_SLog.Mustln("    SETOUTMASK_help")
}

func (this *SCmd) GETOUTMASK(vTInputParams bsn_common.TInputParams) {
	this.M_SLog.Mustln(this.M_SLog.M_u32OutMask)
}

func (this *SCmd) GETOUTMASK_help(vTInputParams bsn_common.TInputParams) {
	this.M_SLog.Mustln("    GETOUTMASK_help")
}

func (this *SCmd) INFO(vTInputParams bsn_common.TInputParams) {
	this.M_SLog.Mustln(this.M_SLog)
}

func (this *SCmd) INFO_help(vTInputParams bsn_common.TInputParams) {
	this.M_SLog.Mustln("    INFO_help")
}
