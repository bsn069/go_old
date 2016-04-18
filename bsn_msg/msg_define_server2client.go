package bsn_msg

import (
	"github.com/bsn069/go/bsn_common"
)

const (
	// gate server to client msg
	GMsgDefine_Gate2Client_Min bsn_common.TMsgType = iota

	GMsgDefine_Gate2Client_Ping
	GMsgDefine_Gate2Client_Pong

	GMsgDefine_Gate2Client_Max
)

const (
	// echo server to client msg
	GMsgDefine_Echo2Client_Min bsn_common.TMsgType = iota + GMsgDefine_Gate2Client_Max

	GMsgDefine_Echo2Client_Ping
	GMsgDefine_Echo2Client_Pong

	GMsgDefine_Echo2Client_Max
)
