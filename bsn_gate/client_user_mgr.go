package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

type SClientUserMgr struct {
	*SUserMgr
}

func NewClientUserMgr(vSGate *SGate) (*SClientUserMgr, error) {
	GSLog.Debugln("newClientUserMgr")
	this := &SClientUserMgr{}

	var err error
	this.SUserMgr, err = NewUserMgr(vSGate)
	if err != nil {
		return nil, err
	}

	return this, nil
}

func (this *SClientUserMgr) Close() error {
	if this.M_bClose {
		return errors.New("had close")
	}
	this.M_bClose = true
	this.StopListen()
	return nil
}

func (this *SClientUserMgr) SendMsgToUser(vTUserId bsn_common.TGateUserId, byData []byte) error {
	if len(byData) < int(bsn_msg.CSMsgHeader_Size) {
		return errors.New("too short")
	}

	vIUser, err := this.User(vTUserId)
	if err != nil {
		GSLog.Errorln(err)
		return err
	}
	vSClientUser, _ := vIUser.(*SClientUser)
	vSClientUser.Send(byData)
	return nil
}

func (this *SClientUserMgr) SendMsgToGroup(vTGroupId bsn_common.TGateUserId, byData []byte) error {
	if len(byData) < int(bsn_msg.CSMsgHeader_Size) {
		return errors.New("too short")
	}

	return nil
}

func (this *SClientUserMgr) Run() error {
	go this.runImp()
	return nil
}

func (this *SClientUserMgr) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.Close()
		for _, vIUser := range this.M_TId2User {
			vSClientUser, _ := vIUser.(*SClientUser)
			vSClientUser.Close()
		}
	}()

	for !this.M_bClose {
		vConn, ok := <-this.SListen.M_chanConn
		if !ok {
			GSLog.Errorln("!ok")
			return
		}

		vSUser, err := NewClientUser(this)
		if err != nil {
			GSLog.Errorln(err)
			vConn.Close()
			return
		}

		vSUser.SetConn(vConn)
		vUserId, _ := this.GenUserId()
		vSUser.SetId(vUserId)
		this.AddUser(vSUser)

		vSUser.Run()
	}
}
