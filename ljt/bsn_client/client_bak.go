package bsn_client

// import (
// 	"github.com/bsn069/go/bsn_common"
// 	"github.com/bsn069/go/bsn_input"
// 	"github.com/bsn069/go/bsn_msg"
// 	"github.com/bsn069/go/bsn_net"
// 	// "github.com/bsn069/go/bsn_log"
// 	"errors"
// 	"net"
// 	"strconv"
// 	"sync"
// )

// type TClientId uint16

// var GClientId TClientId = 0

// type SClient struct {
// 	M_TClientId       TClientId
// 	M_strAddr         string
// 	M_Conn            net.Conn
// 	M_bRun            bool
// 	M_Mutex           sync.Mutex
// 	M_chanNotifyClose chan bool
// 	M_chanWaitClose   chan bool
// 	M_SNetConnecter   bsn_net.SNetConnecter
// 	M_SMsgHeader      bsn_msg.SMsgHeader
// }

// func NewClient() (*SClient, error) {
// 	GClientId++
// 	GSLog.Debugln("NewClient() GClientId=", GClientId)

// 	this := &SClient{
// 		M_TClientId:       GClientId,
// 		M_chanNotifyClose: make(chan bool, 0),
// 		M_chanWaitClose:   make(chan bool, 0),
// 	}

// 	vSCmdClient := &SCmdClient{M_SClient: this}
// 	bsn_input.GSInput.Reg("Client"+strconv.Itoa(int(GClientId)), vSCmdClient)

// 	return this, nil
// }

// func (this *SClient) ShowInfo() {
// 	GSLog.Mustln("id :", this.M_TClientId)
// }

// func (this *SClient) SetGateAddr(strAddr string) error {
// 	this.M_strAddr = strAddr
// 	return nil
// }

// func (this *SClient) Conn() net.Conn {
// 	return this.M_Conn
// }

// func (this *SClient) Run() error {
// 	this.M_Mutex.Lock()
// 	defer this.M_Mutex.Unlock()

// 	if this.M_bRun {
// 		return errors.New("running")
// 	}

// 	err := this.Connect()
// 	if err != nil {
// 		return err
// 	}

// 	go this.runImp()
// 	this.M_bRun = true
// 	return nil
// }

// func (this *SClient) Close() error {
// 	this.M_Mutex.Lock()
// 	defer this.M_Mutex.Unlock()

// 	if !this.M_bRun {
// 		return errors.New("not running")
// 	}
// 	GSLog.Mustln("Close begin")

// 	this.M_Conn.Close()
// 	// clear close chanel
// 	select {
// 	case <-this.M_chanNotifyClose:
// 	default:
// 	}
// 	this.M_chanNotifyClose <- true
// 	// wait close complete
// 	<-this.M_chanWaitClose

// 	GSLog.Mustln("Close end")
// 	return nil
// }

// func (this *SClient) Connect() (err error) {
// 	if "" == this.M_strAddr {
// 		return errors.New("no addr")
// 	}

// 	this.M_Conn, err = net.Dial("tcp", this.M_strAddr)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (this *SClient) SendString(strMsg string) error {
// 	return this.SendMsg(bsn_common.TMsgType(0), []byte(strMsg))
// }

// func (this *SClient) Send(byData []byte) error {
// 	vLen := len(byData)
// 	if vLen <= 0 {
// 		return nil
// 	}

// 	writeLen, err := this.Conn().Write(byData)
// 	if err != nil {
// 		return err
// 	}
// 	if writeLen != vLen {
// 		return errors.New("not write all data")
// 	}
// 	return nil
// }

// func (this *SClient) SendMsg(vTMsgType bsn_common.TMsgType, byMsg []byte) error {
// 	vLen := len(byMsg)
// 	if vLen <= 0 {
// 		return nil
// 	}

// 	this.M_SMsgHeader.Fill(vTMsgType, bsn_common.TMsgLen(vLen))
// 	byMsgHeader := this.M_SMsgHeader.Serialize()

// 	err := this.Send(byMsgHeader)
// 	if err != nil {
// 		return err
// 	}

// 	err = this.Send(byMsg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (this *SClient) Recv() error {
// 	err, vu16MsgType, byData := bsn_msg.ReadMsgWithMsgHeader(this.M_Conn)
// 	if err != nil {
// 		if err.Error() == "EOF" {
// 			GSLog.Errorln("connect disconnect")
// 		} else {
// 			GSLog.Errorln("ReadMsg error: ", err)
// 		}
// 		return err
// 	}
// 	GSLog.Debugln("recv msg: ", vu16MsgType, byData)

// 	return err
// }

// func (this *SClient) runImp() {
// 	defer bsn_common.FuncGuard()
// 	defer func() {
// 		GSLog.Debugln("on closing")
// 		this.M_bRun = false

// 		GSLog.Debugln("close connect")
// 		this.M_Conn.Close()

// 		GSLog.Debugln("send notify to wait close chan")
// 		select {
// 		case <-this.M_chanWaitClose:
// 		default:
// 		}
// 		this.M_chanWaitClose <- true

// 		GSLog.Debugln("close ok")
// 	}()

// 	for {
// 		err := this.Recv()
// 		if err != nil {
// 			GSLog.Errorln(err)
// 			break
// 		}
// 	}
// }
