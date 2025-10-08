class Solution {
    fun pacificAtlantic(heights: Array<IntArray>): List<List<Int>> {
        val m = heights.size
        val n = heights[0].size
        val pacific = Array(m) { BooleanArray(n) }
        val atlantic = Array(m) { BooleanArray(n) }
        
        for (i in 0 until m) {
            bfs(heights, pacific, i, 0)
            bfs(heights, atlantic, i, n - 1)
        }

        for (j in 0 until n) {
            bfs(heights, pacific, 0, j)
            bfs(heights, atlantic, m - 1, j)
        }

        val res = mutableListOf<List<Int>>()
        for (i in 0 until m) {
            for (j in 0 until n) {
                if (pacific[i][j] && atlantic[i][j]) {
                    res.add(listOf(i, j))
                }
            }
        }

        return res
    }

    private fun bfs(heights: Array<IntArray>, reachable: Array<BooleanArray>, row: Int, col: Int) {
        val m = heights.size
        val n = heights[0].size

        val queue = mutableListOf(row to col)
        reachable[row][col] = true

        while (queue.isNotEmpty()) {
            val current = queue.removeFirst()
            val currentHeight = heights[current.first][current.second]

            for (step in STEPS) {
                val r = current.first + step.first
                if (r < 0 || r >= m) continue
                
                val c = current.second + step.second
                if (c < 0 || c >= n) continue
                
                if (reachable[r][c] || heights[r][c] < currentHeight) continue
                
                reachable[r][c] = true
                queue.add(Pair(r, c))
            }
        }
    }

    companion object {
        val STEPS = listOf(Pair(-1, 0), Pair(0, -1), Pair(1, 0), Pair(0, 1))
    }
}