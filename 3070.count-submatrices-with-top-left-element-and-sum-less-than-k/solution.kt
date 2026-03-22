class Solution {
    fun countSubmatrices(grid: Array<IntArray>, k: Int): Int {
        val m = grid.size
        val n = grid[0].size
        val sum = IntArray(n)
        var res = 0

        for (i in 0 until m) {
            var p = 0
            for (j in 0 until n) {
                p += grid[i][j]
                sum[j] += p
                if (sum[j] > k) break
                res++
            }
            if (sum[0] > k) break
        }

        return res
    }
}