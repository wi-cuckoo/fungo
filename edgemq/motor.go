package main

import (
	"bytes"
	_ "embed"
	"fmt"
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
	e.ch = make(chan byte, 1)

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
	connID := fmt.Sprintf("%s-%s", conn.RemoteAddr(), conn.LocalAddr())
	log.Printf("accept connection %s\n", connID)
	defer func() {
		log.Printf("close connection %s\n", connID)
		conn.Close()
	}()

	if c, ok := conn.(*net.TCPConn); ok {
		log.Printf("[%s]set keepalive\n", connID)
		c.SetKeepAlive(true)
		c.SetKeepAlivePeriod(time.Second * 2)
	}

	// connect authorization
	passwd := make([]byte, 4)
	if _, err := conn.Read(passwd); err != nil {
		log.Printf("[%s]conn read err: %s", connID, err.Error())
		return
	}
	if !bytes.Equal(passwd, []byte("ojbk")) {
		log.Printf("[%s]conn authorize fail, bad passwd: %s", connID, string(passwd))
		return
	}

	var last byte = 'S'
	t := time.NewTicker(time.Millisecond * 100)
	for range t.C {
		var c byte
		select {
		case c = <-e.ch:
			last = c
		default:
			c = last
			if last == 'S' {
				continue
			}
		}

		log.Printf("[%s]write command: %c\n", connID, c)
		if _, err := conn.Write([]byte{c}); err != nil {
			log.Printf("[%s]got err: %s, close connection\n", connID, err.Error())
			break
		}

	}
}
