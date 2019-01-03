'''
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例: 

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
'''

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        vals = []
        queue = [root]
        while len(queue) > 0:
            l = len(queue)
            for i in range(0, l):
                node = queue[i]
                if node is None:
                    vals.append(None)
                    continue
                vals.append(node.val)
                queue.append(node.left)
                queue.append(node.right)
            queue = queue[l:]
        return "{}".format(vals)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        vals = eval(data)
        if vals[0] is None:
            return None
        vals = [TreeNode(x) if x is not None else None for x in vals]
        root = vals[0]
        l = 1
        while len(vals) > 0: # 有些过于复杂了
            i = l
            for n in range(0, l):
                vals[n].left = vals[i]
                vals[n].right = vals[i+1]
                i += 2
            vals = vals[l:]
            l *= 2
            head = [x for x in vals[:l] if x is not None]
            vals = head + vals[l:]
            l = len(head)
        return root

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))

if __name__ == "__main__":
    tree = TreeNode(1)