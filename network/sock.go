package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const sock = "/tmp/iw.sock"

var notify = make(chan bool)

func init() {
	if _, err := os.Stat(sock); err == nil {
		os.Remove(sock)
	}
}

func main() {

	addr := net.UnixAddr{
		Name: sock,
		Net:  "unix",
	}
	go server(&addr)

	<-notify
	con, err := net.DialUnix(addr.Network(), nil, &addr)
	if err != nil {
		panic(err.Error())
	}
	defer con.Close()

	buf := make([]byte, 512)
	for {
		_, err := con.Write([]byte("sb, who are you"))
		if err != nil {
			fmt.Println("write err: ", err.Error())
		}
		time.Sleep(time.Second * 1)
		size, err := con.Read(buf)
		if err != nil {
			fmt.Println("resp err: ", err.Error())
			continue
		}
		fmt.Println("got resp: ", string(buf[:size]))
	}
}

func server(addr *net.UnixAddr) {
	ln, err := net.ListenUnix(addr.Network(), addr)
	if err != nil {
		panic(err.Error())
	}
	notify <- true

	for {
		con, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go serverCon(con)
	}
}

func serverCon(con net.Conn) {
	buf := make([]byte, 512)
	for {
		size, err := con.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err.Error())
		}
		fmt.Println("read msg: ", string(buf[:size]))
		con.Write([]byte("I'm sb too"))
	}
}
