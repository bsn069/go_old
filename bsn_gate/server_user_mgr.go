package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "time""
	"net"
)

type SServerUserMgr struct {
	*SUserMgr
}

func NewServerUserMgr(vSGate *SGate) (*SServerUserMgr, error) {
	GSLog.Debugln("newServerUserMgr")
	this := &SServerUserMgr{}

	var err error
	this.SUserMgr, err = NewUserMgr(vSGate, this)
	if err != nil {
		return nil, err
	}
	this.SUserMgr.M_TGateFuncOnNewUser = OnNewServerUser

	return this, nil
}

func OnNewServerUser(vImp bsn_common.TVoid, vConn net.Conn) error {
	vSServerUserMgr, ok := vImp.(*SServerUserMgr)
	if !ok {
		return errors.New("OnNewServerUser !ok")
	}
	return vSServerUserMgr.OnNewServerUser(vConn)
}

func (this *SServerUserMgr) OnNewServerUser(vConn net.Conn) error {
	vSUser, err := NewServerUser(this)
	if err != nil {
		return err
	}
	vSUser.SetConn(vConn)
	vSUser.Run()
	return nil
}
