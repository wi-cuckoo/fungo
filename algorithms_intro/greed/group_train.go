package greed

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type rrange struct {
	start, end int
}

// GroupTrain ...
func GroupTrain() {
	buf := []byte("aacbbdcd")
	table := [26]*rrange{}
	for i, b := range buf {
		ch := int(b - 'a')
		if table[ch] == nil {
			table[ch] = &rrange{i, i}
			continue
		}
		table[ch].end = i
	}
	start, end := -1, -1
	for i, b := range buf {
		ch := int(b - 'a')
		if start == -1 {
			start, end = i, table[ch].end
		}
		if end == i {
			fmt.Println(i - start + 1)
			start = -1
			continue
		}
		if table[ch].end > end {
			end = table[ch].end
		}
	}
}
