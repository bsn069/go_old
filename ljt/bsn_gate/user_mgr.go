package bsn_gate

import (
	"github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "net"
	"errors"
	"net"
	"sync"
)

type TGateFuncOnNewUser func(vImp bsn_common.TVoid, vConn net.Conn) error
type SUserMgr struct {
	M_TUserMgrType bsn_common.TGateUserMgrType
	M_TUserId      bsn_common.TGateUserId
	M_TId2User     bsn_common.TGateUserId2User
	*bsn_net.SListen
	M_bRun               bool
	M_SGate              *SGate
	M_chanNotifyClose    chan bool
	M_chanWaitClose      chan bool
	M_Mutex              sync.Mutex
	M_TGateFuncOnNewUser TGateFuncOnNewUser
	M_Imp                bsn_common.TVoid
}

func NewUserMgr(vSGate *SGate, vImp bsn_common.TVoid) (this *SUserMgr, err error) {
	GSLog.Debugln("newUserMgr")
	this = &SUserMgr{
		M_TId2User:        make(bsn_common.TGateUserId2User),
		M_TUserId:         0,
		M_SGate:           vSGate,
		M_bRun:            false,
		M_chanNotifyClose: make(chan bool, 1),
		M_chanWaitClose:   make(chan bool, 1),
		M_Imp:             vImp,
	}
	this.SListen = bsn_net.NewListen()
	return this, nil
}

func (this *SUserMgr) ShowInfo() {
	GSLog.Mustln("listen addr: ", this.M_strAddr)
	GSLog.Mustln("max user id: ", this.M_TUserId)
	GSLog.Mustln("user count : ", len(this.M_TId2User))
	GSLog.Mustln("is running : ", this.M_bRun)
	GSLog.Mustln("is listen  : ", this.IsListen())
}

func (this *SUserMgr) Gate() *SGate {
	return this.M_SGate
}

func (this *SUserMgr) Type() bsn_common.TGateUserMgrType {
	return this.M_TUserMgrType
}

func (this *SUserMgr) GenUserId() (bsn_common.TGateUserId, error) {
	this.M_TUserId++
	return this.M_TUserId, nil
}

func (this *SUserMgr) AddUser(vIUser bsn_common.IGateUser) error {
	this.M_TId2User[vIUser.Id()] = vIUser
	return nil
}

func (this *SUserMgr) DelUser(vIUser bsn_common.IGateUser) error {
	delete(this.M_TId2User, vIUser.Id())
	return nil
}

func (this *SUserMgr) User(vTUserId bsn_common.TGateUserId) (bsn_common.IGateUser, error) {
	return this.M_TId2User[vTUserId], nil
}

func (this *SUserMgr) SendMsgToUser(vTUserId bsn_common.TGateUserId, byData []byte) error {
	if len(byData) < int(bsn_msg.CSMsgHeader_Size) {
		return errors.New("too short")
	}

	vIUser, err := this.User(vTUserId)
	if err != nil {
		GSLog.Errorln(err)
		return err
	}
	return vIUser.Send(byData)
}

func (this *SUserMgr) Close() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if !this.M_bRun {
		return errors.New("not running")
	}
	GSLog.Mustln("Close begin")

	this.StopListen()
	// clear close chanel
	select {
	case <-this.M_chanNotifyClose:
	default:
	}
	this.M_chanNotifyClose <- true
	// wait close complete
	<-this.M_chanWaitClose

	GSLog.Mustln("Close end")
	return nil
}

func (this *SUserMgr) runImp() {
	defer bsn_common.FuncGuard()
	defer func() {
		GSLog.Debugln("on closing")
		this.M_bRun = false

		GSLog.Debugln("close listen")
		this.StopListen()

		GSLog.Debugln("close all not process connect")
		var vbQuit bool = false
		for !vbQuit {
			select {
			case conn := <-this.SListen.M_chanConn:
				conn.Close()
			default:
				vbQuit = true
			}
		}

		GSLog.Debugln("close all user")
		for _, vIUser := range this.M_TId2User {
			vIUser.Close()
		}

		GSLog.Debugln("send notify to wait close chan")
		select {
		case <-this.M_chanWaitClose:
		default:
		}
		this.M_chanWaitClose <- true

		GSLog.Debugln("close all ok")
	}()

	GSLog.Debugln("run imp")
	var vbQuit bool = false
	for !vbQuit {
		vbQuit = true
		select {
		case vConn, ok := <-this.SListen.M_chanConn:
			if !ok {
				GSLog.Errorln("!ok")
				break
			}
			err := this.M_TGateFuncOnNewUser(this.M_Imp, vConn)
			if err != nil {
				GSLog.Errorln(err)
				vConn.Close()
				break
			}

			vbQuit = false
		case <-this.M_chanNotifyClose:
			GSLog.Mustln("receive a notify to close")
		}
	}
}

func (this *SUserMgr) Run() error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	if this.M_bRun {
		return errors.New("running")
	}

	this.Listen()
	go this.runImp()
	this.M_bRun = true
	return nil
}
