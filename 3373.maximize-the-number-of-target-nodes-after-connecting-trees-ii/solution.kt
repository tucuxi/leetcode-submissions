class Solution {
    fun maxTargetNodes(edges1: Array<IntArray>, edges2: Array<IntArray>): IntArray {
        return listOf(edges1, edges2).map { edges -> 
            val adj = adjacency(edges)
            val m = IntArray(adj.size)
            val c = IntArray(3)

            fun dfs(u: Int, o: Int) {
                m[u] = o
                c[o]++
                adj[u]
                    .filter { v -> m[v] < 1 }
                    .forEach { v -> dfs(v, 3 - o) }
            }

            dfs(0, 1)

            c to m
        }
        .let { (cm1, cm2) ->
            val cMax = cm2.first.max()
            val (c1, m1) = cm1
            IntArray(m1.size) { c1[m1[it]] + cMax }
        }
    }

    fun adjacency(edges: Array<IntArray>): Array<MutableSet<Int>> {
        val n = edges.maxOf { (u, v) -> maxOf(u, v) } + 1
        val adj = Array(n) { mutableSetOf<Int>() }
        edges.forEach { (u, v) ->
            adj[u] += v
            adj[v] += u
        }
        return adj
    }
}