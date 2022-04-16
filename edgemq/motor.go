package main

import (
	"bytes"
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
	pub, sub net.Listener
}

func NewEdged(pubAddr, subAddr string) *Edged {
	e := new(Edged)
	var err error
	if e.pub, err = net.Listen("tcp", pubAddr); err != nil {
		panic(err)
	}
	if e.sub, err = net.Listen("tcp", subAddr); err != nil {
		panic(err)
	}
	e.ch = make(chan byte, 8)

	return e
}

func (e *Edged) Serve() {
	go e.serveSubStream()

	r := gin.Default()
	r.PUT("/pub/motor/:cmd", func(c *gin.Context) {
		cmd := c.Param("cmd")
		if len(cmd) < 1 {
			c.String(http.StatusBadRequest, "no cmd")
			return
		}
		for i := 0; i < len(cmd); i++ {
			select {
			case e.ch <- cmd[i]:
			default:
			}
		}
		c.String(http.StatusOK, "ojbk")
	})
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", homePage)
	})

	panic(r.RunListener(e.pub))
}

func (e *Edged) serveSubStream() {
	for {
		conn, err := e.sub.Accept()
		if err != nil {
			log.Printf("[edged] accept conn fail: %s\n", err.Error())
			continue
		}
		go e.handleSubConn(conn)
	}
}

func (e *Edged) handleSubConn(conn net.Conn) {
	log.Printf("accept connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
	defer func() {
		log.Printf("close connection %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
		conn.Close()
	}()

	if c, ok := conn.(*net.TCPConn); ok {
		c.SetKeepAlive(true)
		c.SetKeepAlivePeriod(time.Second * 8)
	}

	// connect authorization
	passwd := make([]byte, 4)
	if _, err := conn.Read(passwd); err != nil {
		log.Printf("conn read err: %s", err.Error())
		return
	}
	if !bytes.Equal(passwd, []byte("ojbk")) {
		log.Printf("conn authorize fail, bad passwd: %s", string(passwd))
		return
	}

	for {
		var c byte
		select {
		case c = <-e.ch:
		case <-time.After(time.Second * 10):
			c = 'S'
		}

		log.Printf("write command: %c\n", c)
		if _, err := conn.Write([]byte{c}); err != nil {
			log.Printf("got err: %s, close connection\n", err.Error())
			break
		}
	}
}
