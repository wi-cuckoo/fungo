package pubsub

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

func TestPubsub(t *testing.T) {
	var (
		wg          sync.WaitGroup
		count       = 100
		testTopic   = "test"
		testMessage = Message{
			Data: []byte("test"),
		}
	)
	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()

	broker := New()
	broker.Create(testTopic)

	sub1, unsub1, _ := broker.Subscribe(testTopic)
	defer unsub1()
	wg.Add(1)
	go func() {
		var cnt int
		for {
			m, ok := <-sub1
			if !ok {
				return
			}
			cnt++
			fmt.Println("sub1 receive msg", cnt, string(m.Data))
			if cnt == count {
				break
			}
		}
		wg.Done()
	}()
	sub2, unsub2, _ := broker.Subscribe(testTopic)
	defer unsub2()
	wg.Add(1)
	go func() {
		var cnt int
		for {
			m, ok := <-sub2
			if !ok {
				return
			}
			cnt++
			fmt.Println("sub2 receive msg", cnt, string(m.Data))
			if cnt == count {
				break
			}
		}
		wg.Done()
	}()

	for i := 0; i < count; i++ {
		if err := broker.Publish(ctx, testTopic, testMessage); err != nil {
			t.Error(err)
		}
	}

	wg.Wait()
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var s = strings.Repeat("a", 1024)

func test() {
	b := []byte(s)
	_ = string(b)
}

func test2() {
	b := str2bytes(s)
	_ = bytes2str(b)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
