import kotlin.math.abs

class Solution {
    fun totalNQueens(n: Int): Int {
        var res = 0
        val queens = IntArray(n)
        
        fun safe(row: Int, col: Int): Boolean =
            (0 until row)
                .none { i ->
                    abs(col - queens[i]).let { it == 0 || it == row - i }
            }

        fun rec(k: Int) {
            if (k == n) {
                res++
            } else {
                for (i in 0 until n) {
                    if (safe(k, i)) {
                        queens[k] = i
                        rec(k + 1)
                    }
                }
            }
        }
        
        rec(0)
        return res
    }
}