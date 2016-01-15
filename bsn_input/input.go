package bsn_input

import (
	"bufio"
	"errors"
	"github.com/bsn069/go/bsn_common"
	// "log"
	"os"
	// "reflect"
	"strings"
)

type tMod2Chan map[string]chan []string

type sInput struct {
	m_tMod2Chan          tMod2Chan
	m_strUseMod          string // use mod name
	m_cmd                *SCmd
	m_quit               bool
	m_tUpperName2RegName map[string]string
	m_bRuning            bool
}

func (this *sInput) run() {
	if this.m_bRuning {
		GLog.Mustln("had run")
		return
	}
	this.m_bRuning = true

	defer this.close()
	for !this.m_quit {
		this.runCmd()
	}
}

func (this *sInput) close() {
	GLog.Mustln("input close")
	this.m_bRuning = false
	this.m_quit = false
}

func (this *sInput) runCmd() {
	defer func() {
		this.m_cmd.m_bShowHelp = false
		if err := recover(); err != nil {
			GLog.Errorln("func return error ", err)
		}
	}()

	r := bufio.NewReader(os.Stdin)

	if this.m_strUseMod == "" {
		GLog.Must(">")
	} else {
		GLog.Must(this.m_strUseMod, ">")
	}

	b, _, _ := r.ReadLine()
	line := string(b)
	if line == "" {
		return
	}

	tokens := strings.Fields(line)
	if len(tokens) < 1 {
		return
	}

	this.m_cmd.m_bShowHelp = false
	strModUpper := strings.ToUpper(tokens[0])
	if this.m_strUseMod == "" || strModUpper == "CD" {
		err := bsn_common.CallStructFunc(this.m_cmd, tokens[0], tokens[1:])
		if err == nil {
			return
		}
	}

	if modChan, ok := this.m_tMod2Chan[this.m_strUseMod]; ok {
		select {
		case modChan <- tokens[1:]:
		default:
			GLog.Mustln("send cmd fail")
		}
		return
	}

	if modChan, ok := this.m_tMod2Chan[strModUpper]; ok {
		select {
		case modChan <- tokens[1:]:
		default:
			GLog.Mustln("send cmd fail")
		}
		return
	}

	err := bsn_common.CallStructFunc(this.m_cmd, "help", tokens[1:])
	if err == nil {
		return
	}
}

func (this *sInput) setUseMod(strMod string) {
	strModUpper := strings.ToUpper(strMod)
	if strModName, ok := this.m_tUpperName2RegName[strModUpper]; ok {
		this.m_strUseMod = strModName
	} else {
		GLog.Errorln("unknonwn mod ", strMod)
	}
}

func (this *sInput) Reg(strMod string) (chan []string, error) {
	strModUpper := strings.ToUpper(strMod)
	if _, ok := this.m_tUpperName2RegName[strModUpper]; ok {
		GLog.Errorln("had reg mod ", strMod)
		return nil, errors.New("had reg mod")
	}

	this.m_tUpperName2RegName[strModUpper] = strMod

	var vChan chan []string = make(chan []string)
	this.m_tMod2Chan[strModUpper] = vChan
	return vChan, nil
}

func (this *sInput) quit() {
	this.m_quit = true
}
