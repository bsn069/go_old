package bsn_common

import (
	// "bsn/bsn_common"
	"time"
)

func SleepSec(vuSec uint) {
	select {
	case <-time.After(time.Duration(time.Second * time.Duration(vuSec))):
	}
}
