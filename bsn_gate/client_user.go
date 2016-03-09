package bsn_gate

import (
	// "errors"
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"unsafe"
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
			break
		}
		GSLog.Debugln("recv this.M_byMsgHeader= ", this.M_byMsgHeader)
		vMsg.DeSerialize(this.M_byMsgHeader)
		GSLog.Debugln("recv vMsg= ", vMsg)

		vLen := vMsg.Len()
		if vLen < bsn_msg.CSMsgHeader_Size {
			GSLog.Errorln("error size")
			break
		}

		this.M_byMsgBody = make([]byte, vLen)
		err = this.ReadMsgBody()
		if err != nil {
			GSLog.Errorln(err)
			return
		}

		GSLog.Debugln("recv this.M_byMsgBody= ", this.M_byMsgBody)
		if vMsg.UserId() == 0 {
			vSMsgHeader := bsn_msg.NewMsgHeaderFromBytes(this.M_byMsgBody)
			GSLog.Debugln("recv vSMsgHeader= ", vSMsgHeader)
			GSLog.Debugln("recv msg= ", string(this.M_byMsgBody[int(unsafe.Sizeof(*vSMsgHeader)):]))
		} else {
			err = vSServerUserMgr.SendMsgToUser(vMsg.UserId(), this.M_byMsgBody)
			if err != nil {
				GSLog.Errorln("process msg err: ", vMsg.Type(), err)
			}
		}
	}
}
