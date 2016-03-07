package bsn_gate

import (
	"errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
)

type SClientUser struct {
	*SUser
	M_SClientUserMgr *SClientUserMgr
}

func NewClientUser(vSClientUserMgr *SClientUserMgr) (*SClientUser, error) {
	GSLog.Debugln("newClientUser")
	this := &SClientUser{}

	vSUser, err := NewUser()
	if err != nil {
		return nil, err
	}
	this.SUser = vSUser
	this.M_SClientUserMgr = vSClientUserMgr

	this.M_byMsgHeader = make([]byte, bsn_msg.CSMsgHeaderGateClient_Size)

	return this, nil
}

func (this *SClientUser) UserMgr() *SClientUserMgr {
	return this.M_SClientUserMgr
}

func (this *SClientUser) Run() {
	go this.runImp()
}

func (this *SClientUser) Close() error {
	if this.M_bClose {
		return errors.New("had close")
	}
	this.M_bClose = true
	this.Conn().Close()
	return nil
}

func (this *SClientUser) Send(byData []byte) error {
	this.Conn().Write(byData)
	return nil
}

func (this *SClientUser) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		this.Close()
		this.UserMgr().DelUser(this)
	}()

	vSServerUserMgr := this.UserMgr().Gate().GetServerMgr()
	var vMsg bsn_msg.SMsgHeaderGateClient
	for !this.M_bClose {
		err := this.ReadMsgHeader()
		if err != nil {
			GSLog.Errorln(err)
			return
		}
		vMsg.DeSerialize(this.M_byMsgHeader)

		vLen := vMsg.Len()
		this.M_byMsgBody = make([]byte, vLen)
		if vLen > 0 {
			err = this.ReadMsgBody()
			if err != nil {
				GSLog.Errorln(err)
				return
			}
		}

		err = vSServerUserMgr.SendMsgToUser(vMsg.UserId(), this.M_byMsgBody)
		if err != nil {
			GSLog.Errorln("process msg err: ", vMsg.Type(), err)
		}
	}
}
