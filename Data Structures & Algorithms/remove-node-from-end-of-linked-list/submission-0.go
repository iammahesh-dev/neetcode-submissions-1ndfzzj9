/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}

	r := l - n 
	if r == 0 {
		return head.Next
	}
	cur = head
	for i:=0; i<r-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}
