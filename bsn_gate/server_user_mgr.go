package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
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

	return this, nil
}

func (this *SServerUserMgr) Close() error {
	if this.M_bClose {
		return errors.New("had close")
	}
	this.M_bClose = true
	this.StopListen()
	return nil
}

func (this *SServerUserMgr) SendMsgToUser(vTUserId bsn_common.TGateUserId, byData []byte) error {
	if len(byData) < int(bsn_msg.CSMsgHeader_Size) {
		return errors.New("too short")
	}

	vIUser, err := this.User(vTUserId)
	if err != nil {
		GSLog.Errorln(err)
		return err
	}
	vSServerUser, _ := vIUser.(*SServerUser)
	return vSServerUser.Send(byData)
}

func (this *SServerUserMgr) SendMsgToGroup(vTGroupId bsn_common.TGateUserId, byData []byte) error {
	if len(byData) < int(bsn_msg.CSMsgHeader_Size) {
		return errors.New("too short")
	}

	return nil
}

func (this *SServerUserMgr) Run() error {
	go this.runImp()
	return nil
}

func (this *SServerUserMgr) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.Close()
		for _, vIUser := range this.M_TId2User {
			vSServerUser, _ := vIUser.(*SServerUser)
			vSServerUser.Close()
		}
	}()

	for !this.M_bClose {
		vConn, ok := <-this.SListen.M_chanConn
		if !ok {
			GSLog.Errorln("!ok")
			return
		}

		vSUser, err := NewServerUser(this)
		if err != nil {
			GSLog.Errorln(err)
			vConn.Close()
			return
		}

		vSUser.SetConn(vConn)
		vSUser.Run()
	}
}
