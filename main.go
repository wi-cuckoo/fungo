package main

import (
	"fmt"
	"os"
	"time"
)

var pid int

func init() {
	pid = os.Getpid()
}

func main() {
	select {
	case <-time.After(time.Second * 5):
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}
