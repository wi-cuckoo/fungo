package model

import (
	"testing"
)

var tree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 9,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 17,
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
}

func TestBFS(t *testing.T) {
	arr := tree.BFS()
	if len(arr) != 5 {
		t.Fatal("length of tree should be 5")
	}
	t.Log(arr)
}

func TestDFS(t *testing.T) {
	arr := tree.DFS()
	if len(arr) != 5 {
		t.Fatal("length of tree should be 5")
	}
	t.Log(arr)
}

func TestPreOrderTraversal(t *testing.T) {
	nodes := tree.PreOrderTraversal()
	if len(nodes) != 5 {
		t.Fatal("length of tree should be 5")
	}
	t.Log(nodes)
}

func TestInOrderTraversal(t *testing.T) {
	nodes := tree.InOrderTraversal()
	if len(nodes) != 5 {
		t.Fatal("length of tree should be 5")
	}
	t.Log(nodes)
}
