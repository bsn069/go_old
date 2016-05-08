/*
Package bsn_echo.
*/
package bsn_echo

import (
	"bsn_define"
	"github.com/bsn069/go/bsn_log"
)

var (
	GAppName    = "echo"
	GServerType = bsn_define.EServerType_ServerType_Echo
)

var GSLog *bsn_log.SLog
