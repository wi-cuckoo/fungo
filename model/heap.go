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
func (h *Heap) less(i, j int) bool {
	return h.data[i] < h.data[j]
}

// Swap two ele
func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push add an ele
func (h *Heap) Push(x int) {
	h.data = append(h.data, x)
	h.up(h.Len() - 1)
}

// Pop remove and return an ele
func (h *Heap) Pop() int {
	x := h.data[0] // 提取最小值，小顶堆的情况
	n := h.Len() - 1
	h.swap(0, n)         // 把最后一个元素替换根元素成为新的根
	h.down(0, n)         // 下沉改根元素
	h.data = h.data[0:n] //
	return x
}

// List return data
func (h *Heap) List() []int {
	return h.data
}

// NewHeap ...
func NewHeap(nums []int) *Heap {
	// heapify
	h := &Heap{
		data: nums,
	}
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child and j2 is right child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		// now j is the min(left child, right child) if j2 is not the last ele
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

func (h *Heap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}
