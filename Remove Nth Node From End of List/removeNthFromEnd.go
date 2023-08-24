/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // my beautiful solution lol, beats 100% on runtime but somehow doesnt win on memory
 // using a pointer trailing another pointer current by (n+1), we advance current to the end of the linked list
 // just need to make sure that when the trailing pointer has not moved beyond head we deal with it accordingly
 // i.e. when our trailing pointer has not moved beyond the first element in the linked list, then the element we want to remove is certainly the first
 // otherwise the element to be removed is certainly not the first
 // the new element to be linked may be nil or another *ListNode
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    idx := 0
    nPlus1thFromEnd := head
    current := head

    for {
        if current != nil {
            if idx > n {
                nPlus1thFromEnd = nPlus1thFromEnd.Next
            }
            idx++
        }
        if current.Next == nil {
            if idx <= n {
                head = nPlus1thFromEnd.Next
            } else {
                nPlus1thFromEnd.Next = nPlus1thFromEnd.Next.Next
            }
            break
        }
        current = current.Next
    }
    return head
}

// equivalent and a bit better to read
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    idx := 0
    nPlus1thFromEnd := head
    current := head

    for current != nil {
        if idx > n {
            nPlus1thFromEnd = nPlus1thFromEnd.Next
        }
        idx++
        current = current.Next
    }

    if idx <= n {
        head = nPlus1thFromEnd.Next
    } else {
        nPlus1thFromEnd.Next = nPlus1thFromEnd.Next.Next
    }

    return head
}
