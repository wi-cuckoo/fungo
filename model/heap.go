package model

// Heap data structure
type Heap struct {
	data []int
}

// Len return length of ele
func (h *Heap) Len() int {
	return len(h.data)
}

// Less compare two num
func (h *Heap) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}

// Swap two ele
func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push add an ele
func (h *Heap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

// Pop remove and return an ele
func (h *Heap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[0 : h.Len()-1]
	return x
}

// NewHeap ...
func NewHeap(nums []int) *Heap {
	return &Heap{
		data: nums,
	}
}
