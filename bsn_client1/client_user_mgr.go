package bsn_client1

import (
	"errors"
	// "github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_msg"
	"github.com/bsn069/go/bsn_net"
	// "time"
	// "math"
	"net"
	"sync"
)

type TId2ClientUser map[TClientId]*SClientUser
type SClientUserMgr struct {
	*bsn_net.SNetServer

	M_SUserMgr *SUserMgr

	M_MutexUser sync.Mutex
	M_TId2User  TId2ClientUser
	M_TClientId TClientId
}

func NewSClientUserMgr(vSUserMgr *SUserMgr) (*SClientUserMgr, error) {
	GSLog.Debugln("NewSClientUserMgr")
	this := &SClientUserMgr{
		M_SUserMgr:  vSUserMgr,
		M_TClientId: 0,
		M_TId2User:  make(TId2ClientUser, 0),
	}
	this.SNetServer, _ = bsn_net.NewSNetServer(this)

	return this, nil
}

func (this *SClientUserMgr) Run() (err error) {
	return nil
}

func (this *SClientUserMgr) Close() (err error) {
	return nil
}

func (this *SClientUserMgr) UserMgr() *SUserMgr {
	return this.M_SUserMgr
}

func (this *SClientUserMgr) ShowInfo() {
	GSLog.Mustln("max user id: ", this.M_TClientId)
	GSLog.Mustln("user count : ", len(this.M_TId2User))
}

func (this *SClientUserMgr) Send(vTClientId TClientId, vbyMsg []byte) error {
	return nil
}

func (this *SClientUserMgr) NetServerImpAccept(vConn net.Conn) error {
	vSClientUser, err := NewSClientUser(this)
	if err != nil {
		return err
	}

	vTClientId := this.genClientId()
	if vTClientId == 0 {
		return errors.New("genClientId fail")
	}

	vSClientUser.SetConn(vConn)
	vSClientUser.SetId(vTClientId)
	this.addClient(vSClientUser)

	vSClientUser.Run()
	return nil
}

func (this *SClientUserMgr) NetServerImpOnClose() error {
	return nil
}

func (this *SClientUserMgr) addClient(vSClientUser *SClientUser) error {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	this.M_TId2User[vSClientUser.Id()] = vSClientUser
	return nil
}

func (this *SClientUserMgr) delClient(vTClientId TClientId) error {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	delete(this.M_TId2User, vTClientId)
	return nil
}

func (this *SClientUserMgr) getClient(vTClientId TClientId) *SClientUser {
	this.M_MutexUser.Lock()
	defer this.M_MutexUser.Unlock()

	return this.M_TId2User[vTClientId]
}

// generate clientid
// if not generate return 0
func (this *SClientUserMgr) genClientId() TClientId {
	for i := 0; i < 100; i++ {
		this.M_TClientId++
		if this.M_TClientId == 0 {
			continue
		}

		vSClientUser := this.getClient(this.M_TClientId)
		if vSClientUser == nil {
			return this.M_TClientId
		}
	}
	return 0
}
