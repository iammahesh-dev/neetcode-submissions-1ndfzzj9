func findOrder(numCourses int, prerequisites [][]int) []int {
   indeg := make([]int, numCourses)
   adj := make([][]int, numCourses) 
   for i:=0; i < numCourses; i++ {
	adj[i] = make([]int, 0)
   }
   for _, p := range prerequisites {
	src, dst := p[0],p[1]
	indeg[dst]++
	adj[src] = append(adj[src], dst)
   }
   q := make([]int, 0)
   for i,v := range indeg {
	if v==0 {
		q = append(q, i)
	}
   }
   finish := 0
   output := make([]int, 0)
   for len(q)> 0 {
	u := q[0]
	q = q[1:]
	finish++
	output = append(output, u)
	for _, n := range adj[u] {
		indeg[n]--
		if indeg[n] == 0 {
			q = append(q, n)
		}
	}
   }

   if finish != numCourses {
	return []int{}
   }
   for i, j:=0, len(output) - 1; i < j; i,j = i + 1, j - 1 {
	output[i], output[j] = output[j], output[i]
   }
   return output
}
