/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	head := dummy
	L, R := list1, list2
	for L != nil && R != nil {
		if L.Val <= R.Val {
			dummy.Next = L
			L = L.Next
		} else {
			dummy.Next = R
			R = R.Next
		}
		dummy = dummy.Next
	}
	if L != nil {
		dummy.Next = L
	}
	if R != nil {
		dummy.Next = R
	}
	return head.Next
}
