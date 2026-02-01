class Solution {
    fun minTime(n: Int, edges: Array<IntArray>, k: Int): Int {        
        var l = 0
        var r = 1_000_000_000

        while (l < r) {
            val t = l + (r - l) / 2
            val uf = UnionFind(n)

            edges.filter { it[2] > t }.forEach { uf.union(it[0], it[1]) }

            if (uf.count() < k) {
                l = t + 1
            } else {
                r = t
            }
        }

        return l
    }
}

class UnionFind(n: Int) {
    private val parent = IntArray(n) { it }
    private val rank = IntArray(n)

    fun union(x: Int, y: Int) {
        val x = find(x)
        val y = find(y)
        when {
            x == y -> return
            rank[x] < rank[y] -> parent[x] = y
            rank[x] > rank[y] -> parent[y] = x
            else -> {
                parent[x] = y
                rank[y]++
            }
        }
    }

    fun find(x: Int): Int {
        return if (x != parent[x]) {
            find(parent[x]).also { parent[x] = it }
        } else {
            x
        }
    }

    fun count(): Int {
        return parent.withIndex().count { (i, p) -> i == p }
    }
}