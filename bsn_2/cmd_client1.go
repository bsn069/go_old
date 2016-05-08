package bsn_2

// import (
// 	// "errors"
// 	bsn_app "github.com/bsn069/go/bsn_client1"
// 	"github.com/bsn069/go/bsn_common"
// 	// "net"
// 	"strconv"
// 	// "sync"
// 	// "time"
// 	// "os"
// )

// type SCmdClient1 struct {
// 	M_TId2App map[uint32]*bsn_app.SApp
// 	M_AppId   uint32
// }

// func NewCmdClient1() *SCmdClient1 {
// 	this := &SCmdClient1{
// 		M_TId2App: make(map[uint32]*bsn_app.SApp),
// 		M_AppId:   0,
// 	}
// 	return this
// }

// func (this *SCmdClient1) CLIENT1_RUN(vTInputParams bsn_common.TInputParams) {
// 	if len(vTInputParams) != 1 {
// 		GSLog.Errorln("appid")
// 		return
// 	}

// 	vuAppId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vSApp, err := this.Client1Create(uint32(vuAppId))
// 	if err != nil {
// 		GSLog.Errorln(err)
// 		return
// 	}

// 	vSApp.Run()
// }

// func (this *SCmdClient1) Client1Create(vAppId uint32) (*bsn_app.SApp, error) {
// 	vSApp, err := bsn_app.NewSApp(vAppId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	this.M_TId2App[vAppId] = vSApp
// 	return vSApp, nil
// }
