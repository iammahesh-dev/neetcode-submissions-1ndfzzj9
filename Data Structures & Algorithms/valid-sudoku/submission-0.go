func isValidSudoku(board [][]byte) bool {
	// check rows
	for row:=0; row < 9; row++ {
		seen := make(map[byte]bool)
		for i:=0; i < 9; i++{
			if board[row][i] == '.' {
				continue
			}
			if seen[board[row][i]] {
				return false
			}
			seen[board[row][i]] = true
		}
	}
	for col:=0; col < 9; col++ {
		seen := make(map[byte]bool)
		for i:=0; i < 9; i++{
			if board[i][col] == '.' {
				continue
			}
			if seen[board[i][col]] {
				return false
			}
			seen[board[i][col]] = true
		}
	}

	for sqr:=0; sqr < 9; sqr++{
		seen := make(map[byte]bool)
		for i:=0; i<3;i++{
			for j:=0; j<3;j++{
				row := (sqr / 3) * 3 + i
				col := (sqr % 3) * 3 + j
				if board[row][col] == '.'{
					continue
				}
				if seen[board[row][col]] {
					return false
				}
				seen[board[row][col]] = true
			}
		}
	}

	return true
}
