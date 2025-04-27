class Solution {
    companion object {
        val dir = listOf(0 to -1, 0 to 1, 1 to 0, -1 to 0)
    }
    fun longestIncreasingPath(matrix: Array<IntArray>): Int {
        val rows = matrix.size
        val cols = matrix[0].size
        val indegree = Array(rows) { IntArray(cols) }
        
        (0 until rows).forEach { x->
            (0 until cols).forEach { y->
                dir.forEach { (dx, dy) ->
                    val nx = x + dx
                    val ny = y + dy
                    if (nx in 0 until rows && ny in 0 until cols) {
                        if (matrix[nx][ny] < matrix[x][y]) {
                            indegree[x][y]++
                        }
                    }
                }
            }
        }
        
        val queue = ArrayDeque<Pair<Int, Int>>()
        (0 until rows).forEach { x->
            (0 until cols).forEach { y->
                if (indegree[x][y] == 0) {
                    queue.offer(x to y)
                }
            }
        }
        
        var len = 0
        while (queue.isNotEmpty()) {
            len++
            repeat(queue.size) {
                val (x, y) = queue.poll()
                dir.forEach { (dx, dy) ->
                    val nx = x + dx
                    val ny = y + dy
                    if (nx in 0 until rows && ny in 0 until cols) {
                        if (matrix[nx][ny] > matrix[x][y]) {
                            if (--indegree[nx][ny] == 0) {
                                queue.offer(nx to ny)
                            }
                        }
                    }
                }
            }
        }
        return len
    }
}
