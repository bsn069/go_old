package bsn_gate

import (
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_input"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "github.com/bsn069/go/bsn_log"
	// "errors"
	// "net"
	// "strconv"
	// "sync"
	// "bsn_msg_gate_server"
	// "github.com/golang/protobuf/proto"
	// "bsn_define"
	// "time"
)

type SServerUser struct {
	*bsn_net.SConnecterWithMsgHeader
	M_SServerUserMgr *SServerUserMgr
}

func NewSServerUser(vSServerUserMgr *SServerUserMgr) (this *SServerUser, err error) {
	GSLog.Debugln("NewSServerUser")
	this = &SServerUser{
		M_SServerUserMgr: vSServerUserMgr,
	}

	return
}
