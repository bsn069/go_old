package bsn_input

import (
	"bufio"
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "log"
	"os"
	// "reflect"
	// "bytes"
	// "path"
	// "runtime"
	"strings"
	"sync"
	"time"
)

var (
	ErrHadRegMod  = errors.New("had reg mod")
	ErrUnknownMod = errors.New("unknown mod")
	ErrHadRun     = errors.New("had run")
)

type SInput struct {
	M_Mutex                     sync.Mutex
	M_TInputUpperName2RegName   bsn_common.TInputUpperName2RegName
	M_TInputUpperName2CmdStruct bsn_common.TInputUpperName2CmdStruct

	M_strUseMod string // use mod name
	M_SCmd      *SCmd

	M_bQuit   bool
	M_bRuning bool

	M_TInputHelp bsn_common.TInputHelp
}

// set current mod
func (this *SInput) SetUseMod(strMod string) error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	strModUpper := strings.ToUpper(strMod)
	if strModName, ok := this.M_TInputUpperName2RegName[strModUpper]; ok {
		this.M_strUseMod = strModName
	} else {
		GSLog.Errorln(strMod)
		return ErrUnknownMod
	}
	return nil
}

// reg module
func (this *SInput) Reg(strMod string, vICmd interface{}) error {
	this.M_Mutex.Lock()
	defer this.M_Mutex.Unlock()

	strModUpper := strings.ToUpper(strMod)
	if _, ok := this.M_TInputUpperName2RegName[strModUpper]; ok {
		GSLog.Errorln(strMod)
		return ErrHadRegMod
	}
	this.M_TInputUpperName2RegName[strModUpper] = strMod

	this.M_TInputUpperName2CmdStruct[strModUpper] = vICmd

	return nil
}

// call in bin file, don`t in lib will block
func (this *SInput) Run() error {
	this.M_Mutex.Lock()
	if this.M_bRuning {
		this.M_Mutex.Unlock()
		return ErrHadRun
	}
	this.M_bRuning = true
	this.M_Mutex.Unlock()

	for !this.M_bQuit {
		this.runCmd()
	}
	this.M_bRuning = false
	return nil
}

// quit
func (this *SInput) Quit() {
	GSLog.Mustln("input Quit")
	this.M_bQuit = true
}

//
func (this *SInput) ShowFunc(vICmd interface{}) error {
	strsFuncs := bsn_common.Funcs(vICmd)
	for _, strFunc := range strsFuncs {
		if !bsn_common.StringIsUpper(&strFunc) {
			continue
		}

		GSLog.Mustln(strings.ToLower(strFunc))
	}
	return nil
}

func (this *SInput) runCmd() {
	defer GSLog.FuncGuard()

	r := bufio.NewReader(os.Stdin)

	if this.M_strUseMod == "" {
		GSLog.Mustln(">")
	} else {
		GSLog.Mustln(this.M_strUseMod, ">")
	}

	b, _, _ := r.ReadLine()
	line := string(b)
	if line == "" {
		return
	}

	vTInputOK := make(bsn_common.TInputOK, 1)

	go this.doCmd(vTInputOK, line)

	select {
	case <-vTInputOK:
	case <-time.After(time.Duration(2 * time.Second)):
	}
}

func (this *SInput) doCmd(vTInputOK bsn_common.TInputOK, line string) {
	defer func() {
		vTInputOK <- true
		// close(vTInputOK)
	}()

	tokens := strings.Fields(line)
	if len(tokens) < 1 {
		return
	}

	strModUpper := strings.ToUpper(tokens[0])

	var err error
	// show help
	if strModUpper == "?" || strModUpper == "H" || strModUpper == "HELP" {
		vICmd, ok := this.M_TInputUpperName2CmdStruct[strings.ToUpper(this.M_strUseMod)]
		if !ok {
			vICmd = this.M_SCmd
		}

		if len(tokens) < 2 {
			this.ShowFunc(vICmd)
			return
		}

		err = bsn_common.CallStructFunc(vICmd, strings.ToUpper(tokens[1])+"_help", tokens[2:])
		if err != nil {
			GSLog.Errorln("unknown cmd ", tokens[1])
		}
		return
	}

	// system cmd
	err = bsn_common.CallStructFunc(this.M_SCmd, strModUpper, tokens[1:])
	if err == nil {
		return
	}

	// mod cmd
	vICmd, ok := this.M_TInputUpperName2CmdStruct[strings.ToUpper(this.M_strUseMod)]
	if ok {
		err = bsn_common.CallStructFunc(vICmd, strModUpper, tokens[1:])
		if err != nil {
			// GSLog.Errorln(err)
			GSLog.Errorln("unknown cmd ", tokens[0])
		}
		return
	}

	vICmd, ok = this.M_TInputUpperName2CmdStruct[strModUpper]
	if ok {
		err = bsn_common.CallStructFunc(vICmd, strings.ToUpper(tokens[1]), tokens[2:])
		if err != nil {
			// GSLog.Errorln(err)
			GSLog.Errorln("unknown cmd ", tokens[0])
		}
		return
	}

	GSLog.Errorln("unknown mod ", tokens[0])
}
