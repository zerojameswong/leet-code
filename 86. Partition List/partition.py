# beats 51%, a slight change could be to add an empty node as head and just go next then the if not lt_head and if not gte_head blocks are not necessary

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        current = head
        lt_head = lt_current = None
        gte_head = gte_current = None

        while current:
            if current.val < x:
                if not lt_head:
                    lt_head = ListNode(current.val)
                    lt_current = lt_head
                else:
                    next_node = ListNode(current.val)
                    lt_current.next = next_node
                    lt_current = next_node
            else:
                if not gte_head:
                    gte_head = ListNode(current.val)
                    gte_current = gte_head
                else:
                    next_node = ListNode(current.val)
                    gte_current.next = next_node
                    gte_current = next_node

            current = current.next

        if lt_head and gte_head:
            lt_current.next = gte_head
        elif not lt_head and gte_head:
            lt_head = gte_head

        return lt_head
