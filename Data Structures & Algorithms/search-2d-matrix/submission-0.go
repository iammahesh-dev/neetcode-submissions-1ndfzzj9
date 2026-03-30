func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows * cols - 1
	for l <=r {
		mid := l + (r -l) / 2
		row, col := mid / cols, mid % cols
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
