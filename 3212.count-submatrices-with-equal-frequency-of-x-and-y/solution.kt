class Solution {
    fun numberOfSubmatrices(grid: Array<CharArray>): Int {
        val n = grid[0].size
        val px = IntArray(n)
        val py = IntArray(n)
        var res = 0

        for (row in grid) {
            var x = 0
            var y = 0
            for (j in row.indices) {
                when (row[j]) {
                    'X' -> x++
                    'Y' -> y++
                }
                px[j] += x
                py[j] += y
                if (px[j] > 0 && px[j] == py[j]) {
                    res++
                }
            }
        }
        return res
    }
}