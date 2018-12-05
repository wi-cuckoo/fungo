#! /usr/bin/env python

# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        if head is None:
            return False
        p1, p2 = head, head.next
        while p1 != p2:
            if p2 is None or p2.next is None:
                return False
            p1, p2 = p1.next, p2.next.next
        return True