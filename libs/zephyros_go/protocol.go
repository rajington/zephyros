package zephyros_go

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"strconv"
	"strings"
	"io"
	"os"
)


func connect() net.Conn {
	conn, err := net.Dial("tcp", "localhost:1235")
	if err != nil {
		fmt.Println("Can't connect. Is Zephyros running?")
		os.Exit(1)
	}
	return conn
}

var c net.Conn = connect()

var respChans = make(map[float64]chan []byte)

func ListenForCallbacks() {
	reader := bufio.NewReader(c)
	for {
		numBytes, _ := reader.ReadString('\n')
		numBytes = strings.Trim(numBytes, "\n")
		// fmt.Println("bytes", numBytes)
		i, _ := strconv.ParseUint(numBytes, 10, 64)

		buf := make([]byte, i)
		io.ReadFull(reader, buf)

		// fmt.Printf("%#v\n", string(buf))

		var msg []interface{}
		json.Unmarshal(buf, &msg)

		// fmt.Println(msg)

		id, obj := msg[0].(float64), msg[1]
		bytes, _ := json.Marshal(obj)
		respChans[id] <- bytes
	}
}


var msgidChan chan float64 = make(chan float64)

func init() {
	go func() {
		var i float64 = 0
		for {
			i++
			msgidChan <- i
		}
	}()
}

func send(recv float64, fn func([]byte), infinite bool, method string, args ...interface{}) []byte {
	msgid := <- msgidChan

	ch := make(chan []byte, 10) // probably enough
	respChans[msgid] = ch

	msg := []interface{}{msgid, recv, method}
	val, _ := json.Marshal(append(msg, args...))
	jsonstr := string(val)
	fmt.Fprintf(c, "%v\n%v", len(jsonstr), jsonstr)

	if fn == nil {
		resp := <-ch
		delete(respChans, msgid)
		return resp
	}

 	go func() {
		<-ch // ignore

		if infinite {
			for { fn(<-ch) }
		} else {
			fn(<-ch)
		}

		delete(respChans, msgid)
	}()

	return nil
}
