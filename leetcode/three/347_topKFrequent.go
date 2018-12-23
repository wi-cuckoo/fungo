/*
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/

package three

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	cache := map[int]int{}
	for _, v := range nums {
		if _, ok := cache[v]; !ok {
			cache[v] = 1
			continue
		}
		cache[v]++
	}
	h := &FrequentInt{}
	heap.Init(h)
	ck := k
	for i, v := range cache {
		if ck == 0 {
			break
		}
		heap.Push(h, Int{i, v})
		delete(cache, i)
		ck--
	}
	topk := heap.Pop(h).(Int)
	for i, v := range cache {
		if v > topk.Count {
			heap.Push(h, Int{i, v})
			topk = heap.Pop(h).(Int)
		}
	}
	result := make([]int, k)
	result[k-1] = topk.Val
	for h.Len() > 0 {
		k--
		result[k-1] = (heap.Pop(h).(Int)).Val
	}
	return result
}

// Int ...
type Int struct {
	Val   int
	Count int
}

// FrequentInt ...
type FrequentInt []Int

// Len return length of ele
func (f *FrequentInt) Len() int {
	return len(*f)
}

// Less compare two num
func (f *FrequentInt) Less(i, j int) bool {
	return (*f)[i].Count < (*f)[j].Count
}

// Swap two ele
func (f *FrequentInt) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

// Push add an ele
func (f *FrequentInt) Push(x interface{}) {
	*f = append(*f, x.(Int))
}

// Pop remove and return an ele
func (f *FrequentInt) Pop() interface{} {
	x := (*f)[f.Len()-1]
	*f = (*f)[0 : f.Len()-1]
	return x
}
