package network

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
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
	fmt.Println("accept connection from ", conn.RemoteAddr().String())
	// 如果不调用 Close 方法，将不会向客户端发送 FIN
	// 此时如果客户端主动发起关闭，服务端将进入 CLOSE_WAIT 状态，而客户端处于 FIN_WAIT 状态
	// 对于 FIN_WAIT 会有一个超时时间，超时后关闭，而服务端会一直保持，除非跟着进程一起凉凉
	defer conn.Close()

	rb := bufio.NewReader(conn)
	for {
		c, err := rb.ReadByte()
		if err != nil || c != '$' {
			break
		}
		// read int
		b, err := rb.ReadBytes('\r')
		if err != nil {
			break
		}
		fmt.Println("read data len:", b)
		n, _ := ParseInt64(b[:len(b)-1])
		buf := make([]byte, n+1)
		if _, err := io.ReadFull(rb, buf); err != nil {
			break
		}
		fmt.Println("recv ", len(buf), string(buf))
		// conn.Write([]byte("ojbk\n"))
		// break // 主动退出模拟服务端主动关闭链接，即主动发出FIN，收到 ACK 后进入 TIME_WAIT 状态
	}
}

// ParseInt64 ...
func ParseInt64(bs []byte) (int64, error) {
	var n int64
	if len(bs) == 0 || bs[0] == '0' {
		return n, errors.New("invalid byte")
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] < '0' || bs[i] > '9' {
			return n, errors.New("invalid byte")
		}
		n = n*10 + int64(bs[i]-'0')
	}
	return n, nil
}

// Dial TCP Server
func Dial(addr string) {
	con, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	str := "strconv.Atoi(string(b))"
	wb := bufio.NewWriter(con)
	for i := 0; i < 4; i++ {
		wb.WriteByte('$')
		n := len(str) - i
		wb.Write([]byte(strconv.FormatInt(int64(n), 10)))
		wb.WriteByte('\r')
		wb.Write([]byte(str[:n]))
		wb.WriteByte('\r')
		wb.Flush()
	}
	time.Sleep(time.Minute)
}
