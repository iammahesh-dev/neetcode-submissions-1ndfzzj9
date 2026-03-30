/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    temp := head
	for temp != nil {
		copy := &Node{Val: temp.Val}
		copy.Next = temp.Next
		temp.Next = copy
		temp = copy.Next
	}

	temp = head

	for temp != nil {
		if temp.Random != nil {
			copy := temp.Next
			copy.Random = temp.Random.Next
		}
		temp = temp.Next.Next
	}
	temp = head
	dummy := &Node{}
	newList := dummy

	for temp != nil {
		copy := temp.Next
		newList.Next = copy
		newList = copy
		temp.Next = copy.Next
		temp = temp.Next
	}
	return dummy.Next
}
