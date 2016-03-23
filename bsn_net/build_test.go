package bsn_net

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestBase(t *testing.T) {

}

func TestListen(t *testing.T) {
	runtime.GOMAXPROCS(2)
	vIListen := NewListen()
	vIListen.SetListenAddr("localhost:50000")
	vIListen.Listen()

	for i := 0; i < 10; i++ {
		go func(index int) {
			fmt.Println("listen" + strconv.Itoa(index))
			vConn, ok := <-vIListen.M_chanConn
			fmt.Println("listen on " + strconv.Itoa(index))
			if !ok {
				fmt.Println("!ok")
				return
			}
			defer vConn.Close()
			fmt.Println(vConn.RemoteAddr().String())
		}(i)
	}

	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println("Dial" + strconv.Itoa(index))
			vConn, err := net.Dial("tcp", "localhost:50000")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer vConn.Close()

			buf := make([]byte, 1)
			n, err := vConn.Read(buf)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(n)
			}
		}(i)
	}

	<-time.After(time.Second * 1)
	fmt.Println("end")
	vIListen.StopListen()
	fmt.Println("end-------")
}
