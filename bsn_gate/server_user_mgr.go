package bsn_gate

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "time"
	// "net"
	// "bsn_msg_gate_gateconfig"
	// "strconv"
	// "sync"
)

type SServerUserMgr struct {
	*bsn_common.SState
	M_SApp  *SApp
	M_Users []*SServerUser
}

func NewSServerUserMgr(vSApp *SApp) (this *SServerUserMgr, err error) {
	GSLog.Debugln("NewSServerUserMgr")
	this = &SServerUserMgr{
		M_SApp: vSApp,
	}
	this.SState = bsn_common.NewSState()

	return
}

func (this *SServerUserMgr) run() (err error) {
	return
}

func (this *SServerUserMgr) close() (err error) {
	return
}
