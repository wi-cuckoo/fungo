package main

import (
	"bufio"
	_ "embed"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed ace.html
var homePage []byte

func main() {
	e := NewEdged(":4399", ":9999")
	e.Serve()
}

type Edged struct {
	ch       chan byte
	up, down net.Listener
}

func NewEdged(upAddr, downAddr string) *Edged {
	e := new(Edged)
	var err error
	if e.up, err = net.Listen("tcp", upAddr); err != nil {
		panic(err)
	}
	if e.down, err = net.Listen("tcp", downAddr); err != nil {
		panic(err)
	}
	e.ch = make(chan byte, 2)

	return e
}

func (e *Edged) Serve() {
	go e.serveDownStream()

	r := gin.Default()
	r.PUT("/pub/motor/:cmd", func(c *gin.Context) {
		cmd := c.Param("cmd")
		if len(cmd) < 1 {
			c.String(http.StatusBadRequest, "no cmd")
			return
		}
		select {
		case e.ch <- cmd[0]:
			c.String(http.StatusOK, "ojbk")
		default:
			c.String(http.StatusInsufficientStorage, "no capacity")
		}
	})
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", homePage)
	})

	panic(r.RunListener(e.up))
}

func (e *Edged) serveDownStream() {
	for {
		conn, err := e.down.Accept()
		if err != nil {
			log.Printf("[edged] accept conn fail: %s\n", err.Error())
			continue
		}
		go e.handleDownConn(conn)
	}
}

func (e *Edged) handleDownConn(conn net.Conn) {
	log.Printf("accept connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
	defer func() {
		log.Printf("close connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
		conn.Close()
	}()

	if c, ok := conn.(*net.TCPConn); ok {
		c.SetKeepAlive(true)
		c.SetKeepAlivePeriod(time.Second * 8)
	}

	wb := bufio.NewWriter(conn)
	for {
		recv := make([]byte, 4)
		if _, err := conn.Read(recv); err != nil {
			break
		}
		log.Println("recv msg: ", string(recv))

		var c byte
		select {
		case c = <-e.ch:
		case <-time.After(time.Second * 10):
			c = 'S'
		}

		if _, err := wb.Write([]byte{c, '\r', '\n'}); err != nil {
			log.Printf("got err: %s, close connection\n", err.Error())
			break
		}
		wb.Flush()
		log.Println("sent command")
	}
}
