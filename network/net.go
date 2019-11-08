package network

import (
	"fmt"
	"net"
)

// Serve ...
func Serve(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("serve on ", ln.Addr().String())
	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}
		go serveCon(conn)
	}
}

func serveCon(conn net.Conn) {
	// 主动调用 Close 会向客户端发送 FIN，如果不这么做将会导致服务端
	// 进入 Close_Wait 状态，无法自拔
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
