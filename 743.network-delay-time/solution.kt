data class Distance(val node: Int, val dist: Int)

typealias AdjList = Array<MutableList<Distance>>

class Solution {
    fun networkDelayTime(times: Array<IntArray>, n: Int, k: Int): Int {
        val adj = makeAdjList(times, n)
        val dist = bfs(adj, k)
        val max = dist.sliceArray(1..n).max()!!
        return if (max == Int.MAX_VALUE) -1 else max
    }
    
    private fun makeAdjList(edges: Array<IntArray>, n: Int): AdjList {
        return Array(n + 1) { mutableListOf<Distance>() }.apply {
            edges.forEach { (u, v, w) -> this[u].add(Distance(v, w)) }
        }
    }
    
    private fun bfs(adj: AdjList, k: Int): IntArray {
        val dist = IntArray(adj.size + 1) { Int.MAX_VALUE }
        val queue = ArrayDeque<Distance>()
        queue.offer(Distance(k, 0))
        while (queue.isNotEmpty()) {            
            val (u, t) = queue.poll()
            if (t < dist[u]) {
                dist[u] = t
                adj[u].forEach { (v, w) -> queue.offer(Distance(v, t + w)) }
            }
        }
        return dist
    }
}