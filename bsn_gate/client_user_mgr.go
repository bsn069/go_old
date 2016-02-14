package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
)

type SClientUserMgr struct {
	*SUserMgr
}

func newClientUserMgr() (*SClientUserMgr, error) {
	GLog.Debugln("newClientUserMgr")
	this := &SClientUserMgr{}

	var err error
	this.SUserMgr, err = newUserMgr()
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

func (this *SClientUserMgr) SendToUser(vTUserId TUserId, byData []byte) error {
	vIUser := this.GetUser(vTUserId)
	vSClientUser, _ := vIUser.(*SClientUser)
	vSClientUser.Send(byData)
	return nil
}

func (this *SClientUserMgr) SendToGroup(vTGroupId TUserId, byData []byte) error {
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
		for _, vIUser := range this.M_users {
			vSClientUser, _ := vIUser.(*SClientUser)
			vSClientUser.Close()
		}
	}()

	for !this.M_bClose {
		vConn, ok := <-this.SListen.M_chanConn
		if !ok {
			GLog.Errorln("!ok")
			return
		}

		vSUser, err := newClientUser(this)
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
