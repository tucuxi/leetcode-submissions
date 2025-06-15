class Solution {
    fun kthSmallest(matrix: Array<IntArray>, k: Int): Int {
        val pq = PriorityQueue<Pair<Int, Int>>(matrix.size, compareBy { (i, j) -> matrix[i][j] })
        matrix.indices.forEach { pq.add(Pair(0, it)) }
        repeat(k-1) {
            val (i, j) = pq.poll()
            if (i < matrix.lastIndex) {
                pq.offer(Pair(i + 1, j))
            }
        }
        val (i, j) = pq.poll()
        return matrix[i][j]
    }
}