package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
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
	dir := os.TempDir()
	filename := filepath.Join(dir, "write-concurrency")
	texts := []string{text2, text1, text3}
	rand.Shuffle(len(texts), func(i, j int) {
		texts[i], texts[j] = texts[j], texts[i]
	})

	// write concurrently
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			text := texts[idx%3]
			ioutil.WriteFile(filename, []byte(text), os.ModePerm)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// sync write
	// for _, t := range texts {
	// 	ioutil.WriteFile(filename, []byte(t), os.ModePerm)
	// }

	fmt.Println(filename)
}

const text1 = `
import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)`

const text2 = `
module github.com/wi-cuckoo/fungo

go 1.12

require (
	github.com/containerd/cgroups v0.0.0-20190328223300-4994991857f9
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/godbus/dbus v0.0.0-20190413140323-8e900ab0295c // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/hashicorp/consul/api v1.1.0
)`

const text3 = `
    _______           ________     
   / ____(_)_______  / ____/ /_  __
  / /_  / / ___/ _ \/ /_  / / / / /
 / __/ / / /  /  __/ __/ / / /_/ / 
/_/   /_/_/   \___/_/   /_/\__, /  
                          /____/  hello, my lord! `
