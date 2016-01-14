package bsn_common

import (
	// "bsn/bsn_common"
	"errors"
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	"reflect"
	"runtime"
	"strings"
)

func CallStructFunc(sStruct interface{}, strFunc string, strParams []string) error {
	strFuncUpper := strings.ToUpper(strFunc)
	vValue := reflect.ValueOf(sStruct)
	vFunc := vValue.MethodByName(strFuncUpper)
	if vFunc.IsValid() {
		vArgs := []reflect.Value{reflect.ValueOf(strParams)}
		vFunc.Call(vArgs)
	} else {
		return errors.New("unknonwn func " + strFunc)
	}
	return nil
}

func GetCallerFileLine(calldepth int) (file string, line int) {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

func GetCallInfo(calldepth int) (pkgName, funcName, filePath string, line int, err error) {
	pc, filePath, line, ok := runtime.Caller(calldepth)
	if !ok {
		err = errors.New("runtime.Caller fail")
		return
	}
	vFunc := runtime.FuncForPC(pc)
	if vFunc == nil {
		err = errors.New("runtime.FuncForPC fail")
		return
	}

	allFuncName := vFunc.Name()
	index := strings.LastIndex(allFuncName, "/")
	if index != -1 {
		allFuncName = allFuncName[index+1:]
	}

	strArray := strings.SplitN(allFuncName, ".", 2)
	pkgName, funcName = strArray[0], strArray[1]
	return
}
