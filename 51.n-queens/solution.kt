import kotlin.math.abs

class Solution {
    fun solveNQueens(n: Int): List<List<String>> {
        val res = mutableListOf<List<String>>()
        val queens = IntArray(n)
        
        fun add(): Unit = queens
            .map { i -> (0 until n)
                    .map { j -> if (j == i) 'Q' else '.' }
                    .joinToString("")
            }
            .let { res.add(it) }

        
        fun safe(row: Int, col: Int): Boolean =
            (0 until row).none { i -> abs(col - queens[i]).let { it == 0 || it == row - i } }

        fun rec(k: Int): Unit =
            if (k == n) {
                add()
            } else {
                for (i in 0 until n) {
                    if (safe(k, i)) {
                        queens[k] = i
                        rec(k + 1)
                    }
                }
            }
        
        rec(0)
        return res
    }
}