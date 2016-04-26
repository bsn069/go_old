package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	// "time"
	"net"
)

type SClientUserMgr struct {
	*SUserMgr
}

func NewClientUserMgr(vSGate *SGate) (*SClientUserMgr, error) {
	GSLog.Debugln("newClientUserMgr")
	this := &SClientUserMgr{}

	var err error
	this.SUserMgr, err = NewUserMgr(vSGate, this)
	if err != nil {
		return nil, err
	}
	this.SUserMgr.M_TGateFuncOnNewUser = OnNewClientrUser

	return this, nil
}

func OnNewClientrUser(vImp bsn_common.TVoid, vConn net.Conn) error {
	vSClientUserMgr, ok := vImp.(*SClientUserMgr)
	if !ok {
		return errors.New("OnNewClientrUser !ok")
	}

	return vSClientUserMgr.OnNewClientrUser(vConn)
}

func (this *SClientUserMgr) OnNewClientrUser(vConn net.Conn) error {
	vSUser, err := NewClientUser(this)
	if err != nil {
		return err
	}

	vSUser.SetConn(vConn)
	vUserId, _ := this.GenUserId()
	vSUser.SetId(vUserId)
	this.AddUser(vSUser)

	vSUser.Run()
	return nil
}
