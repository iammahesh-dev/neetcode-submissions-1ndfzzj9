/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

	for len(lists) > 1 { 
		mergedList := []*ListNode{}
		for i:=0; i< len(lists); i += 2{
			l1 := lists[i]
			var l2 *ListNode
			if i + 1 < len(lists) {
				l2 = lists[i+1]
			}
			mergedList = append(mergedList, mergeTwoLists(l1, l2)) 
		}
		lists = mergedList
	}
	return lists[0]
}

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