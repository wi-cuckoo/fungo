package pubsub

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRegNotify(t *testing.T) {
	rn := NewRegNotify()

	go func() {
		cnt := 0
		for i := 0; ; i++ {
			err := rn.Issue(context.Background(), "test-portal", fmt.Sprintf("G1 fuck %d", i))
			if err != nil {
				fmt.Println("G1 issue payload err:", err.Error())
				cnt++
			}
			if cnt >= 2 {
				break
			}
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		cnt := 0
		for i := 0; ; i++ {
			err := rn.Issue(context.Background(), "test-portal", fmt.Sprintf("G2 fuck %d", i))
			if err != nil {
				fmt.Println("G2 issue payload err:", err.Error())
				cnt++
			}
			if cnt >= 2 {
				break
			}
		}
		time.Sleep(time.Millisecond)
	}()

	ti := time.NewTimer(time.Second * 2)
	ch, _ := rn.Register(context.Background(), "test-portal")

	if _, err := rn.Register(context.Background(), "test-portal"); err != ErrExisted {
		t.Fatal("can't be register more than once")
	}
	for {
		select {
		case p, ok := <-ch:
			fmt.Println(p, ok)
		case <-ti.C:
			fmt.Println("time to shutdown pipe")
			rn.Unregister(context.Background(), "test-portal")
			<-time.After(time.Second * 3)
			return
		}
	}
}
