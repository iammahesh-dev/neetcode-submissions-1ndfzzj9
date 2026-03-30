func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
	area := 0

	directions := [][]int{{-1,0}, {0, 1}, {1,0}, {0, -1}}

	var bfs func(r, c int) int
	bfs = func(r, c int) int {
		q := [][]int{{r,c}}
		grid[r][c] = 0
		res := 1

		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			for _,d := range directions {
				nr,nc := front[0] + d[0], front[1] + d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == 0{
					continue
				}
				q = append(q, []int{nr, nc})
				grid[nr][nc] = 0
				res++
			} 
		}
		return res
	}

	for i:=0; i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == 1 {
				area = max(area, bfs(i, j))
			}
		}
	}
	return area
}
