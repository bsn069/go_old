package bsn_msg

import (
	"github.com/bsn069/go/bsn_common"
)

const (
	// server to gate msg
	GMsgDefine_Server2Gate_Min bsn_common.TMsgType = iota

	GMsgDefine_Server2Gate_Ping
	GMsgDefine_Server2Gate_Pong

	GMsgDefine_Server2Gate_ClientMsg

	GMsgDefine_Server2Gate_Max
)

const (
	// gate config server to gate msg
	GMsgDefine_GateConfigServer2Gate_Min bsn_common.TMsgType = iota + GMsgDefine_Server2Gate_Max

	GMsgDefine_GateConfig2Gate_Reg

	GMsgDefine_GateConfig2Gate_Max
)
