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
	this.SUserMgr, err = NewUserMgr(vSGate)
	if err != nil {
		return nil, err
	}
	this.SUserMgr.M_TGateFuncOnNewUser = OnNewServerUser

	return this, nil
}

func OnNewServerUser(vIGateUserMgr bsn_common.IGateUserMgr, vConn net.Conn) error {
	vSServerUserMgr, ok := vIGateUserMgr.(*SServerUserMgr)
	if !ok {
		return errors.New("!ok")
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
