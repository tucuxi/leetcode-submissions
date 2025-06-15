class Solution {
    fun shortestPathBinaryMatrix(grid: Array<IntArray>): Int {
        val N = grid.lastIndex
        var length = 0
        val queue = ArrayDeque<Pair<Int, Int>>()
        queue.offer(0 to 0)
        while (queue.isNotEmpty()) {
            length++
            repeat(queue.size) {
                val (r, c) = queue.poll()
                if (grid[r][c] == 0) {
                    grid[r][c] = 2
                    if (r == N && c == N) {
                        return length
                    }
                    for (rr in maxOf(0, r - 1) .. minOf(N, r + 1)) {
                        for (cc in maxOf(0, c - 1) .. minOf(N, c + 1)) {
                            if (grid[rr][cc] == 0) {
                                queue.offer(rr to cc)
                            } 
                        }
                    }
                }
            }
        }
        return -1
    }
}