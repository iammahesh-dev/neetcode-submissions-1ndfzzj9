/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)
	oldToNew[node] = &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	queue := []*Node{node}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, nei := range cur.Neighbors {
			if _,exists := oldToNew[nei]; !exists {
				oldToNew[nei] = &Node{Val: nei.Val, Neighbors: make([]*Node, 0)}
				queue = append(queue, nei)
			}
			oldToNew[cur].Neighbors = append(oldToNew[cur].Neighbors, oldToNew[nei])
		}
	}
	return oldToNew[node]
}
