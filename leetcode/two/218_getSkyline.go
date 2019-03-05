/*
 天际线问题
*/
package two

import (
	"sort"

	"github.com/wi-cuckoo/fungo/model"
)

func getSkyline(buildings [][]int) [][]int {
	dots := make([][]int, len(buildings)*2)
	for i, b := range buildings {
		dots[2*i] = []int{b[0], b[2]}
		dots[2*i+1] = []int{b[1], b[2]}
	}
	sort.Slice(dots, func(i, j int) bool {
		a, b := dots[i], dots[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	// fmt.Println(dots)
	h := model.NewHeap([]int{})
	for _, d := range dots {
		h.Push(d[0])
	}
	return nil
}
