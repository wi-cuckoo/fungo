// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
package one

import (
	"github.com/wi-cuckoo/fungo/model"
)

/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(preorder []int, inorder []int) *model.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	pos := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			pos = i
			break
		}
	}
	//fmt.Println("pos is ", pos, preorder, inorder)
	linorder, rinorder := inorder[:pos], inorder[pos+1:]
	lpreorder, rpreorder := preorder[1:pos+1], preorder[pos+1:]
	root := &model.TreeNode{
		Val:   rootVal,
		Left:  buildTree(lpreorder, linorder),
		Right: buildTree(rpreorder, rinorder),
	}
	return root
}
