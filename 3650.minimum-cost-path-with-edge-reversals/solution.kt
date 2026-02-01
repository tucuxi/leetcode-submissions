class Solution {
    fun minCost(n: Int, edges: Array<IntArray>): Int {
        val graph = Array(n) { mutableListOf<Pair<Int, Int>>() }.also {
            edges.forEach { (u, v, w) ->
                it[u].add(v to w)
                it[v].add(u to 2 * w)
            }
        }

        val distance = IntArray(n) { Int.MAX_VALUE }.also { it[0] = 0 }

        val visited = BooleanArray(n)

        val queue = PriorityQueue<Pair<Int, Int>>(compareBy({ it.second })).also { it.offer(0 to 0) }

        while (queue.isNotEmpty()) {
            val (node, dist) = queue.poll()
            
            if (node == n - 1) {
                return dist
            }
            if (visited[node]) {
                continue
            }
            visited[node] = true

            graph[node].forEach { (v, w) ->
                if (dist + w < distance[v]) {
                    distance[v] = dist + w
                    queue.offer(v to distance[v])
                }
            }
        }

        return -1
    }
}