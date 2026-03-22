class Solution {
    fun minAbsDiff(grid: Array<IntArray>, k: Int): Array<IntArray> {
        val m = grid.size
        val n = grid[0].size
        return Array(m - k + 1) { i ->
            IntArray(n - k + 1) { j ->
                val s = TreeSet<Int>()
                for (r in i until i + k) {
                    for (c in j until j + k) {
                        s += grid[r][c]
                    }
                }
                
                s.windowed(2).minOfOrNull { (a, b) -> b - a } ?: 0
            }
        }
    }
}