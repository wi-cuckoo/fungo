package model

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	h := NewHeap(nums)
	heap.Init(h)
	t.Log(h.List())
	val := heap.Pop(h)
	if val.(int) != 1 {
		t.Error("should be 1, not", val)
	}
}

func TestBFS(t *testing.T)  {
	tree := &TreeNode{
		Val:3,
		Left: &TreeNode{
			Val:9,
		},
		Right:&TreeNode{
			Val:20,
			Left:&TreeNode{
				Val:17,
			},
			Right: &TreeNode{
				Val:5,
			},
		},
	}
	arr := tree.BFS()
	if len(arr) != 5 {
		t.Fatal("length of tree should be 5")
	}
	t.Log(arr)
}

func TestDFS(t *testing.T)  {
	tree := &TreeNode{
		Val:3,
		Left: &TreeNode{
			Val:9,
		},
		Right:&TreeNode{
			Val:20,
			Left:&TreeNode{
				Val:17,
			},
			Right: &TreeNode{
				Val:5,
			},
		},
	}
	arr := tree.DFS()
	if len(arr) != 5 {
		t.Fatal("length of tree should be 5")
	}
	t.Log(arr)
}