package bsn_log

import (
	"fmt"
	"sync"
)

var fl_muOutput sync.Mutex // ensures atomic writes; protects the following fields

func output(strInfo string) {
	// fl_muOutput.Lock()
	// defer fl_muOutput.Unlock()
	fmt.Print(strInfo)
}
