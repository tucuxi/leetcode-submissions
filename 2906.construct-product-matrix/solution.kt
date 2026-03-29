const val MOD = 12345L

class Solution {
    fun constructProductMatrix(grid: Array<IntArray>): Array<IntArray> {
        val n = grid.size
        val m = grid[0].size
        val p = Array(n) { IntArray(m) }

        var suffix = 1L
        for (i in n - 1 downTo 0) {
            for (j in m - 1 downTo 0) {
                p[i][j] = suffix.toInt()
                suffix = suffix * grid[i][j] % MOD
            }
        }

        var prefix = 1L
        for (i in 0 until n) {
            for (j in 0 until m) {
                p[i][j] = (p[i][j] * prefix % MOD).toInt()
                prefix = prefix * grid[i][j] % MOD
            }
        }

        return p
    }
}