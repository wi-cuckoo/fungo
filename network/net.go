package network

import (
	"fmt"
	"net"
	"time"
)

// Serve ...
func Serve(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("serve on ", ln.Addr().String())
	time.Sleep(time.Minute * 5)
	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}
		go serveCon(conn)
	}
}

func serveCon(conn net.Conn) {
	fmt.Println("accept connection from ", conn.RemoteAddr().String())
	// 如果不调用 Close 方法，将不会向客户端发送 FIN
	// 此时如果客户端主动发起关闭，服务端将进入 CLOSE_WAIT 状态，而客户端处于 FIN_WAIT 状态
	// 对于 FIN_WAIT 会有一个超时时间，超时后关闭，而服务端会一直保持，除非跟着进程一起凉凉
	defer conn.Close()
	for {
		buf := make([]byte, 1<<9)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err.Error())
			break
		}
		conn.Write([]byte("ojbk\n"))
		// break // 主动退出模拟服务端主动关闭链接，即主动发出FIN，收到 ACK 后进入 TIME_WAIT 状态
	}
}

// Dial TCP Server
func Dial(addr string) {
	con, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer con.Close()
}
