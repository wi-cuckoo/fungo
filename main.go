package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wi-cuckoo/fungo/pubsub"
)

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
	rn := pubsub.NewRegNotify()

	go func() {
		for i := 0; ; i++ {
			err := rn.Issue(context.Background(), "test-portal", fmt.Sprintf("G1 fuck %d", i))
			if err != nil {
				fmt.Println("G1 issue payload err:", err.Error())
				break
			}
			// time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for i := 0; ; i++ {
			err := rn.Issue(context.Background(), "test-portal", fmt.Sprintf("G2 fuck %d", i))
			if err != nil {
				fmt.Println("G2 issue payload err:", err.Error())
				break
			}
		}
		time.Sleep(time.Millisecond)
	}()

	ti := time.NewTimer(time.Second * 2)
	ch, _ := rn.Register(context.Background(), "test-portal")
	defer rn.Unregister(context.Background(), "test-portal")
	for {
		select {
		case p, ok := <-ch:
			fmt.Println(p, ok)
		case <-ti.C:
			fmt.Println("time to shutdown pipe")
			return
		}
	}
}
