const val MOD = 1_000_000_007L

data class Entry(val r: Int, val c: Int, val p: Long)

class Solution {
    fun maxProductPath(grid: Array<IntArray>): Int {
        val m = grid.size
        val n = grid[0].size
        val gt = Array(m) { LongArray(n) }
        val lt = Array(m) { LongArray(n) }

        gt[0][0] = grid[0][0].toLong()
        lt[0][0] = gt[0][0]

        for (j in 1 until n) {
            gt[0][j] = gt[0][j - 1] * grid[0][j].toLong()
            lt[0][j] = gt[0][j]
        }

        for (i in 1 until m) {
            gt[i][0] = gt[i - 1][0] * grid[i][0].toLong()
            lt[i][0] = gt[i][0]
        }

        for (i in 1 until m) {
            for (j in 1 until n) {
                val maxPrev = maxOf(gt[i][j - 1], gt[i - 1][j])
                val minPrev = minOf(lt[i][j - 1], lt[i - 1][j])
                val g = grid[i][j].toLong()
                if (g >= 0) {
                    gt[i][j] = maxPrev * g
                    lt[i][j] = minPrev * g
                } else {
                    gt[i][j] = minPrev * g
                    lt[i][j] = maxPrev * g
                }
            }
        }

        return gt[m - 1][n - 1].let { if (it < 0) -1 else (it % MOD).toInt() }
    }
}