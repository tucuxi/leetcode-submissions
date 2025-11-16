class Solution {
    fun rangeAddQueries(n: Int, queries: Array<IntArray>): Array<IntArray> {
        val p = Array(n + 1) { IntArray(n + 1) }
        
        queries.forEach { (r1, c1, r2, c2) ->
            p[r1][c1]++
            p[r1][c2 + 1]--
            p[r2 + 1][c1]--
            p[r2 + 1][c2 + 1]++
        }

        val res = Array(n) { IntArray(n) }

        for (i in 0 until n) {
            for (j in 0 until n) {
                val x = listOf(
                    if (i == 0) 0 else res[i - 1][j],
                    if (j == 0) 0 else res[i][j - 1],
                    if (i == 0 || j == 0) 0 else -res[i - 1][j - 1]
                )
                res[i][j] = p[i][j] + x.sum() 
            }
        }

        return res
    }
}