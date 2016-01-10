package bsn_log

import (
	"strconv"
	"testing"
)

func TestPrint(t *testing.T) {
	Print("this is Print1")
	Print("this is Print2")
	Print("this is Print3")
}

func TestPrintln(t *testing.T) {
	Println("this is Println1")
	Println("this is Println2", 1, false)
	Println("this is Println3", t, true, 3)
}

func TestPrintf(t *testing.T) {
	Printf("this is Printf1 %v \n", 1)
	Printf("this is Printf1 %v %v\n", 1, 2)
	Printf("this is Printf1 %v %v %v\n", 1, 2, 3)
}

func TestParal(t *testing.T) {
	syncChan := make(chan bool, 10)
	for i := 0; i < 5; i++ {
		go func(i int) {
			Println("aaaaaaaaaaaaaaaaaaaa" + strconv.Itoa(i))
			syncChan <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		go func(i int) {
			Println("bbbbbbbbbbbbbbbbbbbb" + strconv.Itoa(i))
			syncChan <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-syncChan
	}
}
