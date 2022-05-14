package main

import (
	"bytes"
	_ "embed"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	PART_BOUNDARY       = "123456789000000000000987654321"
	STREAM_CONTENT_TYPE = "multipart/x-mixed-replace;boundary=" + PART_BOUNDARY
	STREAM_BOUNDARY     = "\r\n--" + PART_BOUNDARY + "\r\n"
	STREAM_PART         = "Content-Type: image/jpeg\r\n\r\n"
)

//go:embed ace.html
var homePage []byte

func main() {
	e := NewEdged(":4399", ":9999")
	e.Serve()
}

type Edged struct {
	ch       chan byte
	streamOn bool
	frameCh  chan *Frame
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
	e.streamOn = true
	e.frameCh = make(chan *Frame, 128)
	return e
}

func (e *Edged) Serve() {
	go e.serveSubStream()

	r := gin.Default()
	r.PUT("/pub/cam/:op", func(c *gin.Context) {
		op := c.Param("op")
		switch op {
		case "on":
			e.pushCmd('O')
			e.streamOn = true
		case "off":
			e.pushCmd('C')
			e.streamOn = false
		}
	})
	r.PUT("/pub/motor/:cmd", func(c *gin.Context) {
		cmd := c.Param("cmd")
		if len(cmd) < 1 {
			c.String(http.StatusBadRequest, "no cmd")
			return
		}
		for i := 0; i < len(cmd); i++ {
			e.pushCmd(cmd[i])
		}
		c.String(http.StatusOK, "ojbk")
	})
	r.GET("/stream", func(c *gin.Context) {
		c.Header("Content-Type", STREAM_CONTENT_TYPE)
		for e.streamOn {
			select {
			case <-c.Done():
				return
			case f := <-e.frameCh:
				c.Writer.Write([]byte(STREAM_BOUNDARY))
				c.Writer.Write([]byte(STREAM_PART))
				c.Writer.Write(f.Buf)
			}
		}
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

func (e *Edged) pushCmd(c byte) {
	select {
	case e.ch <- c:
	default:
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

	quit := make(chan struct{})
	go func() {
		for {
			n, err := readFrameLen(conn)
			if err != nil {
				log.Printf("readFrameLen fail: %s\n", err.Error())
				break
			}
			log.Printf("readFrameLen got: %d\n", n)
			frame, err := readFrame(conn, int(n))
			if err != nil {
				log.Printf("readFrame fail: %s\n", err.Error())
				break
			}
			select {
			case e.frameCh <- &Frame{Buf: frame, Len: n}:
			default:
			}

		}
		quit <- struct{}{}
	}()

	for {
		var c byte
		select {
		case c = <-e.ch:
		case <-quit:
			return
		}
		log.Printf("[%s]write command: %c\n", connID, c)
		if _, err := conn.Write([]byte{c}); err != nil {
			log.Printf("[%s]got err: %s, close connection\n", connID, err.Error())
			break
		}
	}
}

type Frame struct {
	Len    uint32 // Length of the buffer in bytes
	Width  uint32 // Width of the buffer in pixels
	Height uint32 // Height of the buffer in pixels
	Buf    []byte // Pointer to the pixel data
	// pixformat_t format; //  Format of the pixel data
	// int64 timestamp;   // Timestamp since boot of the first DMA buffer of the frame
}

func readFrameLen(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	return binary.BigEndian.Uint32(buf), err
}

func readFrame(r io.Reader, n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}
