package model

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS 实现广度优先搜索
func (t *TreeNode) BFS() []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	queue := []*TreeNode{t} // 使用队列来协助遍历
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			nodes = append(nodes, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return nodes
}

// DFS 实现深度优先搜索
func (t *TreeNode) DFS() []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	stack := []*TreeNode{t} // 使用栈来协助遍历
	for len(stack) > 0 {
		// 父节点出栈
		p := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		nodes = append(nodes, p.Val)
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return nodes
}

// PreOrderTraversal 实现前序遍历 - 递归版本
func (t *TreeNode) PreOrderTraversal() []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{t.Val}
	if t.Left != nil {
		nodes = append(nodes, t.Left.PreOrderTraversal()...)
	}
	if t.Right != nil {
		nodes = append(nodes, t.Right.PreOrderTraversal()...)
	}

	return nodes
}

// PosOrderTraversal 实现后序遍历 - 递归版本
func (t *TreeNode) PosOrderTraversal() []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	if t.Left != nil {
		nodes = append(nodes, t.Left.PosOrderTraversal()...)
	}
	if t.Right != nil {
		nodes = append(nodes, t.Right.PosOrderTraversal()...)
	}
	nodes = append(nodes, t.Val)

	return nodes
}

// InOrderTraversal 实现中序遍历 - 递归版本
func (t *TreeNode) InOrderTraversal() []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	if t.Left != nil {
		nodes = append(nodes, t.Left.InOrderTraversal()...)
	}
	nodes = append(nodes, t.Val)
	if t.Right != nil {
		nodes = append(nodes, t.Right.InOrderTraversal()...)
	}
	return nodes
}
