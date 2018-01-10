#meet a test question is 
#Example
#Given 1->2->3->3->4->5->3, val = 3, you should return the list as 1->2->4->5




answer :

"""
Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
"""


class Solution:
    """
    @param: head: a ListNode
    @param: val: An integer
    @return: a ListNode
    """
    def removeElements(self, head, val):
        # write your code here
        if head is None:
            return None
        while head.val ==val:
            head = head.next
            if (head == None):
                return None
        pre = head
        
#赋值之后，pre改变则head也同样会改变

        while pre.next is not None:
            if pre.next.val == val:
                pre.next = pre.next.next
            else:
                pre = pre.next
        return head
