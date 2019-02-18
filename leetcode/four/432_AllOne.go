/*
实现一个数据结构支持以下操作：

Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
挑战：以 O(1) 的时间复杂度实现所有操作。
*/

package four

import "math"

// AllOne ...
type AllOne struct {
	keys map[string]int // 记录每个key对应的值
	// 他们说用双向链表，保存最大最小节点。。。。。
}

// Constructor Initialize your data structure here
func Constructor() AllOne {
	return AllOne{
		keys: make(map[string]int),
	}
}

// Inc ...
func (a *AllOne) Inc(key string) {
	if _, ok := a.keys[key]; ok {
		a.keys[key]++
		return
	}
	a.keys[key] = 1
}

// Dec ...
func (a *AllOne) Dec(key string) {
	v, ok := a.keys[key]
	if !ok {
		return
	}
	if v == 1 {
		delete(a.keys, key)
		return
	}
	a.keys[key]--
}

// GetMaxKey ..
func (a *AllOne) GetMaxKey() string {
	max, key := 0, ""
	for k, v := range a.keys {
		if v > max {
			key, max = k, v
		}
	}
	return key
}

// GetMinKey ...
func (a *AllOne) GetMinKey() string {
	min, key := math.MaxInt32, ""
	for k, v := range a.keys {
		if v < min {
			key, min = k, v
		}
	}
	return key
}
