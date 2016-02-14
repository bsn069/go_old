package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
)

type SServerUserMgr struct {
	*SUserMgr
}

func newServerUserMgr() (*SServerUserMgr, error) {
	GLog.Debugln("newServerUserMgr")
	this := &SServerUserMgr{}

	var err error
	this.SUserMgr, err = newUserMgr()
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

func (this *SServerUserMgr) SendToUser(vTUserId TUserId, byData []byte) error {
	vIUser := this.GetUser(vTUserId)
	vSServerUser, _ := vIUser.(*SServerUser)
	vSServerUser.Send(byData)
	return nil
}

func (this *SServerUserMgr) SendToGroup(vTGroupId TUserId, byData []byte) error {
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
		for _, vIUser := range this.M_users {
			vSServerUser, _ := vIUser.(*SServerUser)
			vSServerUser.Close()
		}
	}()

	for !this.M_bClose {
		vConn, ok := <-this.SListen.M_chanConn
		if !ok {
			GLog.Errorln("!ok")
			return
		}

		vSUser, err := newServerUser(this)
		if err != nil {
			GLog.Errorln(err)
			vConn.Close()
			return
		}

		vSUser.SetConn(vConn)
		vUserId, _ := this.GenId()
		vSUser.SetId(vUserId)
		this.AddUser(vSUser)

		vSUser.Run()
	}
}
