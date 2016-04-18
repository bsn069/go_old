package bsn_msg

import (
	"github.com/bsn069/go/bsn_common"
)

const (
	// client to gate server msg
	GMsgDefine_Client2Gate_Min bsn_common.TMsgType = iota

	GMsgDefine_Client2Gate_Ping
	GMsgDefine_Client2Gate_Pong

	GMsgDefine_Client2Gate_Max
)

const (
	// client to echo server msg
	GMsgDefine_Client2Echo_Min bsn_common.TMsgType = iota + GMsgDefine_Client2Gate_Max

	GMsgDefine_Client2Echo_Ping
	GMsgDefine_Client2Echo_Pong

	GMsgDefine_Client2Echo_Max
)
