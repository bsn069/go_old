package bsn_client1

import (
	"errors"
	// "fmt"
	// "github.com/bsn069/go/bsn_msg"
)

func (this *SClientUser) procMsg() error {
	GSLog.Debugln(this.M_SMsgHeader)
	GSLog.Debugln(this.M_by2MsgBody)

	return errors.New("unknown msg")
}
