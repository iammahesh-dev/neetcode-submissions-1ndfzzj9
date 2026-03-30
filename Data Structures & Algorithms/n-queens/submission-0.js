class Solution {
    /**
     * @param {number} n
     * @return {string[][]}
     */
    solveNQueens(n) {
        const col = new Set()
        const pDiag = new Set()
        const nDiag = new Set()
        const res = []
        const board = Array.from({ length: n }, () => Array.from({ length: n }, () => '.'))
        function backtrack(r) {
            if (r === n) {
                res.push(board.map((b) => b.join("")))
                return
            }
            for(let i = 0; i < n; i++) {
                if(col.has(i) || pDiag.has(r + i) || nDiag.has(r - i)){
                    continue
                }

                col.add(i)
                pDiag.add(r + i)
                nDiag.add(r - i)
                board[r][i] = 'Q'

                backtrack(r + 1)

                col.delete(i)
                pDiag.delete(r + i)
                nDiag.delete(r - i)
                board[r][i] = '.'
            }   
        }
        backtrack(0)
        return res
    }
}
