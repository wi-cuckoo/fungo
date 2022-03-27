package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var cmdCh chan byte = make(chan byte, 10)

func main() {
	http.HandleFunc("/fuck", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		cmd := q.Get("cmd")
		if len(cmd) == 1 {
			cmdCh <- cmd[0]
		}
		w.Write([]byte("OJBK"))
		w.Header().Add("content-type", "text/plain")
	})
	go http.ListenAndServe("192.168.3.39:4399", nil)

	ln, err := net.Listen("tcp", "192.168.3.39:9999")
	if err != nil {
		panic(err)
	}

	fmt.Println("serve on ", ln.Addr().String())
	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}
		fmt.Printf("accept connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
		go serveCon(conn)
	}
}

func serveCon(conn net.Conn) {
	defer func() {
		fmt.Printf("close connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
		conn.Close()
	}()

	if c, ok := conn.(*net.TCPConn); ok {
		fmt.Println("set keepalive")
		c.SetKeepAlive(true)
	}

	wb := bufio.NewWriter(conn)
	for {
		recv := make([]byte, 4)
		if _, err := conn.Read(recv); err != nil {
			break
		}
		fmt.Println("recv msg: ", string(recv))

		var c byte
		select {
		case c = <-cmdCh:
		case <-time.After(time.Second * 10):
			c = 'S'
		}

		if _, err := wb.Write([]byte{c, '\r', '\n'}); err != nil {
			fmt.Printf("got err: %s, close connection\n", err.Error())
			break
		}
		wb.Flush()
		fmt.Println("sent command")
	}
}
