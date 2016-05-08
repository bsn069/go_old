package bsn_gate

import (
	// "errors"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	// "fmt"
	"net"
	// "sync"
)

type SClientUser struct {
	M_SClientUserMgr *SClientUserMgr
	M_Conn           net.Conn
	M_TClientId      uint16
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr, vConn net.Conn, vClientId uint16) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_Conn:           vConn,
		M_TClientId:      vClientId,
	}

	return this, nil
}

func (this *SClientUser) ClientId() uint16 {
	return this.M_TClientId
}

func (this *SClientUser) close() (err error) {
	if this.M_Conn != nil {
		this.M_Conn.Close()
		this.M_Conn = nil
	}
	return
}
