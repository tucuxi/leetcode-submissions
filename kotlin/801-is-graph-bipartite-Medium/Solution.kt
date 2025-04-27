class Solution {
    fun isBipartite(graph: Array<IntArray>): Boolean {
        val dsu = DSU(graph.size)
        for (i in graph.indices) {
            for (j in graph[i]) {
                if (dsu.find(i) == dsu.find(j)) {
                    return false
                }
                dsu.union(graph[i][0], j)
            }
        }
        return true
    }
}

private class DSU(n: Int) {
    private val p = IntArray(n) { it }

    fun union(u: Int, v: Int) {
        p[find(u)] = find(v)
    }
    
    fun find(u: Int): Int {
        if (p[u] != u) {
            p[u] = find(p[u])
        }
        return p[u]
    }
}