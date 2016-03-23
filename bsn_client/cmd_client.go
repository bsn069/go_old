package bsn_client

import (
	"github.com/bsn069/go/bsn_common"
	// "errors"
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	// "math/rand"
	// "reflect"
	// "strconv"
	"strings"
)

type SCmdClient struct {
	M_SClient *SClient
}

func (this *SCmdClient) SHOW_INFO(vTInputParams bsn_common.TInputParams) {
	this.M_SClient.ShowInfo()
}

func (this *SCmdClient) SEND_STRING(vTInputParams bsn_common.TInputParams) {
	if len(vTInputParams) < 1 {
		GSLog.Errorln("msg strings")
		return
	}

	strMsg := strings.Join([]string(vTInputParams), " ")

	err := this.M_SClient.SendString(strMsg)
	if err != nil {
		GSLog.Errorln(err)
	}
}
