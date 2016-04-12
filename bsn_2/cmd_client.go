package bsn_2

import (
	// "errors"
	"github.com/bsn069/go/bsn_client"
	"github.com/bsn069/go/bsn_common"
	// "github.com/bsn069/go/bsn_gate"
	// "net"
	"strconv"
	// "sync"
	// "time"
	// "os"
	"fmt"
)

type SCmdClient struct {
}

func NewCmdClient() *SCmdClient {
	this := &SCmdClient{}
	return this
}

func (this *SCmdClient) CLIENT_RUN(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("gateid")
		return
	}

	vuGateId, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	strAddr := fmt.Sprintf("localhost:%v", 40000+vuGateId)
	vSClient, err := bsn_client.NewClient(strAddr)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	err = vSClient.Run()
	if err != nil {
		GSLog.Errorln(err)
		return
	}
}
