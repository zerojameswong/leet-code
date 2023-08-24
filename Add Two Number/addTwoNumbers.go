/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var resultListHead = new(ListNode)
    var resultListCurrent = resultListHead

    done := false
    var l1Val, l2Val, carry int
    for !done {
        if l1 != nil {
            l1Val = l1.Val
            l1 = l1.Next
        } else {
            l1Val = 0
        }
        if l2 != nil {
            l2Val = l2.Val
            l2 = l2.Next
        } else {
            l2Val = 0
        }

        sum := l1Val + l2Val + carry
        store := sum % 10
        carry = sum / 10
        resultListCurrent.Val = store

        if l1 == nil && l2 == nil && carry == 0 {
            done = true
        } else {
            resultListCurrent.Next = new(ListNode)
            resultListCurrent = resultListCurrent.Next
        }
    }
    return resultListHead
}