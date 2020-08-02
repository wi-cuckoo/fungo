/*
给定一个二叉树，原地将它展开为一个单链表。
例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package one

import "github.com/wi-cuckoo/fungo/model"

//方法一：深度优先遍历，消耗空间
func flatten(root *model.TreeNode) {
	if root == nil {
		return
	}
	dummy := &model.TreeNode{}
	stack := []*model.TreeNode{root} // 使用栈来协助遍历
	for len(stack) > 0 {
		// 父节点出栈
		p := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		dummy.Right = p
		p.Left = nil
		dummy = dummy.Right
	}
}

// 方法二：深度优先遍历，不需要额外空间
func flattenV2(root *model.TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
