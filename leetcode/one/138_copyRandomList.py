'''
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的深度拷贝
'''

# Definition for singly-linked list with a random pointer.
class RandomListNode(object):
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        if head is None:
            return None
        ref = dict()
        index = list()
        p, i = head, 0
        while p is not None:
            node = RandomListNode(p.label)
            index.append(node)
            ref[p] = i
            p, i = p.next, i+1
        p, i = head, 0
        while p is not None:
            if i < len(index)-1:
                index[i].next = index[i+1]
            r = ref.get(p.random, -1)
            if r >= 0:
                index[i].random = index[r]
            p, i = p.next, i+1
        return index[0]

if __name__ == '__main__':
    Solution().copyRandomList(None)