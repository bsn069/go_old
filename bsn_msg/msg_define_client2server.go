package bsn_msg

import (
	"github.com/bsn069/go/bsn_common"
)

const (
	// client to gate server msg
	GMsgDefine_Client2Gate_Min bsn_common.TMsgType = iota
	GMsgDefine_Client2Gate_Ping
	GMsgDefine_Client2Gate_Echo
	GMsgDefine_Client2Gate_Max

	GMsgDefine_Client2Server_Max
)
