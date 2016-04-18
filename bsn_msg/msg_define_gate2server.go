package bsn_msg

import (
	"github.com/bsn069/go/bsn_common"
)

const (
	//  gate to server msg
	GMsgDefine_Gate2Server_Min bsn_common.TMsgType = iota

	GMsgDefine_Gate2Server_Ping
	GMsgDefine_Gate2Server_Pong

	GMsgDefine_Gate2Server_ClientMsg

	GMsgDefine_Gate2Server_Max
)

const (
	//  gate to gate config server msg
	GMsgDefine_Gate2GateConfigServer_Min bsn_common.TMsgType = iota + GMsgDefine_Gate2Server_Max

	GMsgDefine_Gate2GateConfig_Reg

	GMsgDefine_Gate2GateConfig_Max
)
