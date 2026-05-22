class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = len(board)
        cols = len(board)
        for i in range(rows):
            seen = {}
            for j in range(cols):
                if board[i][j] == ".":
                    continue
                if board[i][j] in seen:
                    return False
                seen[board[i][j]] = True
        
        for i in range(rows):
            seen = {}
            for j in range(cols):
                if board[j][i] == ".":
                    continue
                if board[j][i] in seen:
                    return False
                seen[board[j][i]] = True

        for sq in range(9):
            seen = {}
            for i in range(3):
                for j in range(3):
                    row = math.floor(sq / 3) * 3 + i
                    col = sq % 3 * 3 + j
                    if board[row][col] == ".":
                        continue
                    if board[row][col] in seen:
                        return False
                    
                    seen[board[row][col]] = True
        
        return True