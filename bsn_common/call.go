package bsn_common

import (
	// "bsn/bsn_common"
	// "errors"
	// "bufio"
	"fmt"
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
		// fmt.Println("run func ", strFunc)
		vArgs := []reflect.Value{reflect.ValueOf(strParams)}
		vFunc.Call(vArgs)
	} else {
		fmt.Println("unknonwn func ", strFunc)
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
