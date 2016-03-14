package bsn_gate2

import (
	// "errors"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "unsafe"
	"net"
)

type TClientId uint16

type SClientUser struct {
	M_SClientUserMgr *SClientUserMgr
	M_TClientId      TClientId
	M_Conn           net.Conn
}

func NewSClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("NewSClientUser")
	this := &SClientUser{
		M_SClientUserMgr: vSClientUserMgr,
		M_TClientId:      0,
		M_Conn:           nil,
	}

	return this, nil
}

func (this *SClientUser) Run() {

}

func (this *SClientUser) Close() {
	this.M_SClientUserMgr.delClient(this.Id())
	this.M_SClientUserMgr = nil
}

func (this *SClientUser) SetId(vTClientId TClientId) error {
	this.M_TClientId = vTClientId
	return nil
}

func (this *SClientUser) Id() TClientId {
	return this.M_TClientId
}

func (this *SClientUser) SetConn(vConn net.Conn) error {
	this.M_Conn = vConn
	return nil
}
