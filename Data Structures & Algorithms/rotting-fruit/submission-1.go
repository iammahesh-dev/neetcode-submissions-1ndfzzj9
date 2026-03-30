func orangesRotting(grid [][]int) int {
    q := make([][]int, 0)
	rows, cols := len(grid), len(grid[0])

	time, fresh := 0,0

	for i:=0;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				q = append(q, []int{i,j})
			}
		}
	}

	dirs := [][]int{{0,1}, {0, -1}, {1,0}, {-1,0}}

	for len(q) > 0 && fresh > 0 {
		l := len(q)
		for i:=0; i < l; i++ {
			rotten := q[0]
			q = q[1:]
			for _,d := range dirs {
				nr,nc := rotten[0] + d[0], rotten[1] + d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				q = append(q, []int{nr, nc})
				fresh--
			}
		}
		time++
	}

	if fresh != 0 {
		return -1
	}

	return time

}
