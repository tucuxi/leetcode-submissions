class Solution {
    fun minimumFlips(n: Int, edges: Array<IntArray>, start: String, target: String): List<Int> {
        val adj = Array(n) { ArrayList<Pair<Int, Int>>() }
        edges.forEachIndexed { i, (u, v) ->
            adj[u].add(v to i)
            adj[v].add(u to i)
        }

        val diff = IntArray(n) { i ->
            if (start[i] != target[i]) 1 else 0
        }

        val res = mutableListOf<Int>()

        fun dfs(u: Int, p: Int): Boolean {
            var needsFlip = diff[u] == 1
    
            adj[u].forEach { (v, i) ->
                if (v != p && dfs(v, u)) {
                    res.add(i)
                    needsFlip = !needsFlip
                }
            }
            return needsFlip
        }

        return if (dfs(0, -1)) {
            listOf(-1)
        } else {
            res.sort()
            res
        }
    }
}