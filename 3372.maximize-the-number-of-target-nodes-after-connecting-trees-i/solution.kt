class Solution {
    fun maxTargetNodes(edges1: Array<IntArray>, edges2: Array<IntArray>, k: Int): IntArray {
        val adj1 = adjacency(edges1)
        val adj2 = adjacency(edges2)
        val count2 = adj2.indices.maxOf { i -> bfs(i, k - 1, adj2) }
        return IntArray(adj1.size) { i -> bfs(i, k, adj1) + count2 }
    }

    fun adjacency(edges: Array<IntArray>): Array<MutableSet<Int>> {
        val n = edges.maxOf { (u, v) -> maxOf(u, v) } + 1
        val adj = Array(n) { mutableSetOf<Int>() }
        edges.forEach { (u, v) ->
            adj[u].add(v)
            adj[v].add(u)
        }
        return adj
    }

    fun bfs(start: Int, t: Int, adj: Array<MutableSet<Int>>): Int {
        val visited = BooleanArray(adj.size)
        visited[start] = true
        var queue = ArrayDeque<Int>(listOf(start))
        var distance = 0
        var count = 0
        while (queue.isNotEmpty() && distance <= t) {
            val nextLevel = ArrayDeque<Int>()
            while (queue.isNotEmpty()) {
                val u = queue.removeFirst()
                count++
                adj[u].forEach { v ->
                    if (!visited[v]) { 
                        nextLevel.add(v)
                        visited[v] = true
                    }
                }
            }
            queue = nextLevel
            distance++
        }
        return count
    }
}