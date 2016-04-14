package bsn_msg

import (
	"github.com/bsn069/go/bsn_common"
)

const (
	// gate server to client msg
	GMsgDefine_Gate2Client_Min bsn_common.TMsgType = iota
	GMsgDefine_Gate2Client_Pong
	GMsgDefine_Gate2Client_Echo
	GMsgDefine_Gate2Client_Max
)
