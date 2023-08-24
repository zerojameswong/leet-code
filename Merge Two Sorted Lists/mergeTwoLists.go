/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// noice
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var mergedHead, merged *ListNode

    for list1 != nil || list2 != nil {
        if list1 != nil && list2 == nil {
            if merged != nil {
                merged.Next = list1
            } else {
                mergedHead = list1
                merged = mergedHead
            }
            list1 = nil
        } else if list1 == nil && list2 != nil {
            if merged != nil {
                merged.Next = list2
            } else {
                mergedHead = list2
                merged = mergedHead
            }
            list2 = nil
        } else {    
            list1Val := list1.Val
            list2Val := list2.Val
            var minVal int

            if list1Val <= list2Val {
                minVal = list1Val
                list1 = list1.Next
            } else {
                minVal = list2Val
                list2 = list2.Next
            }

            newNode := &ListNode{Val: minVal}
            if merged != nil {
                merged.Next = newNode
                merged = merged.Next
            } else {
                mergedHead = newNode
                merged = mergedHead
            }
        }
    }
    return mergedHead
}
