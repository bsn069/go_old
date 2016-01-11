package main

import (
	"fmt"
	// "github.com/bsn069/go/bsn_common"
	"github.com/bsn069/go/bsn_input"
	"github.com/bsn069/go/bsn_log"
)

// var _ = mImp1.Main

func procInput(strName string, vInputChan chan []string) {
	for {
		vStringArr := <-vInputChan
		fmt.Println("this is ", strName, vStringArr)
	}
}

func main() {
	GLog := bsn_log.New("main")
	GLog.Debugln(1)
	vInputChan1, _ := bsn_input.Instance().Reg("1")
	go procInput("mod1", vInputChan1)

	vInputChan2, _ := bsn_input.Instance().Reg("2")
	go procInput("mod2", vInputChan2)

	bsn_input.Run()
}

/*
gate -id 1 -level 1 -port 40001
gate -id 2 -level 1 -port 40002 -UpNodeAddr localhost:40001
gate -id 3 -level 1 -port 40003 -UpNodeAddr localhost:40001
gate -id 4 -level 1 -port 40004 -UpNodeAddr localhost:40001
*/
