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

func (this *SCmd) CLIENT_RUN(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) != 1 {
		GSLog.Errorln("gateport")
		return
	}

	vuGatePort, err := strconv.ParseUint(vTInputParams[0], 10, 32)
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	vSClient, err := bsn_client.NewClient()
	if err != nil {
		GSLog.Errorln(err)
		return
	}

	strAddr := fmt.Sprintf("localhost:%v", vuGatePort)
	vSClient.SetGateAddr(strAddr)

	err = vSClient.Run()
	if err != nil {
		GSLog.Errorln(err)
		return
	}
}
