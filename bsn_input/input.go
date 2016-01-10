package bsn_input

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bsn069/go/bsn_common"
	// "log"
	"os"
	// "reflect"
	"strings"
)

type tMod2Chan map[string]chan []string

type sInput struct {
	m_strSep             string
	m_tMod2Chan          tMod2Chan
	m_strUseMod          string // use mod name
	m_cmd                *SCmd
	m_quit               bool
	m_tUpperName2RegName map[string]string
	m_bRuning            bool
}

func (this *sInput) run() {
	if this.m_bRuning {
		fmt.Println("had run")
		return
	}
	this.m_bRuning = true

	defer this.close()
	for !this.m_quit {
		this.runCmd()
	}
}

func (this *sInput) close() {
	fmt.Println("input close")
	this.m_bRuning = false
	this.m_quit = false
}

func (this *sInput) runCmd() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("func return error ", err)
		}
	}()

	r := bufio.NewReader(os.Stdin)

	if this.m_strUseMod == "" {
		fmt.Print(">")
	} else {
		fmt.Print(this.m_strUseMod + ">")
	}

	b, _, _ := r.ReadLine()
	line := string(b)

	tokens := strings.Split(line, this.m_strSep)
	// for _, strParam := range tokens {
	// 	fmt.Println(strParam)
	// }

	strMod := tokens[0]
	if strMod == "" {
		this.m_cmd.m_bShowHelp = false
		bsn_common.CallStructFunc(this.m_cmd, tokens[1], tokens[2:])
		this.m_cmd.m_bShowHelp = false
	} else {
		strModUpper := strings.ToUpper(strMod)
		if this.m_strUseMod != "" {
			strModUpper = strings.ToUpper(this.m_strUseMod)
		}

		if modChan, ok := this.m_tMod2Chan[strModUpper]; ok {
			modChan <- tokens[1:]
		} else {
			fmt.Println("unknonwn mod", strMod)
		}
	}

	fmt.Println("")
}

func (this *sInput) setUseMod(strMod string) {
	strModUpper := strings.ToUpper(strMod)
	if strModName, ok := this.m_tUpperName2RegName[strModUpper]; ok {
		this.m_strUseMod = strModName
	} else {
		fmt.Println("unknonwn mod", strMod)
	}
}

func (this *sInput) Reg(strMod string) (chan []string, error) {
	strModUpper := strings.ToUpper(strMod)
	if _, ok := this.m_tUpperName2RegName[strModUpper]; ok {
		fmt.Println("had reg mod ", strMod)
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
