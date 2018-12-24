package model

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 实现广度优先搜索
func (t *TreeNode) BFS() []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	queue := []*TreeNode{t} // 使用堆列来协助遍历
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i <l; i++ {
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

// 实现深度优先搜索
func (t *TreeNode) DFS() []int {
 	if t == nil {
 		return []int{}
	}
	nodes := []int{}
	stack := []*TreeNode{t}
	for len(stack) > 0 {
		// 父节点出栈
		p := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
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
