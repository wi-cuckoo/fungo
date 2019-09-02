package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin" 
)

func init() {
	rand.Seed(time.Now().Unix())
}

func ping() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "failure")
		return
	}
	conn, err := net.DialTimeout("tcp", strings.Join(os.Args[1:], ":"), time.Second*2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Fprintln(os.Stderr, "succ")
	conn.Close()
}
