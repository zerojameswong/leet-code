/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// apparently the faster way to do this is to construct an array from the linked list?
// my way is ugly but its not lazy?
func swapPairs(head *ListNode) *ListNode {
    retHead := head
    
    left := head
    right := head
    for right != nil && right.Next != nil {
        if left == retHead {
            retHead = right.Next
            right.Next = retHead.Next
            retHead.Next = right
            right = right.Next
        } else {
            left.Next = right.Next
            right.Next = left.Next.Next
            left.Next.Next = right
            left = left.Next.Next
            right = right.Next
        }
    }

    return retHead
}
