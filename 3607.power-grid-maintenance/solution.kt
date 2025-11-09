class Solution {
    fun processQueries(c: Int, connections: Array<IntArray>, queries: Array<IntArray>): IntArray {
        val uf = UnionFind(c + 1)

        connections.forEach { (u, v) -> uf.union(u, v) }

        val sets = mutableMapOf<Int, TreeSet<Int>>()

        for (x in 1..c) {
            val p = uf.find(x)
            sets.getOrPut(p) { sortedSetOf<Int>() }.add(x)
        }

        val res = mutableListOf<Int>()

        queries.forEach { (op, x) ->
            val p = uf.find(x)
            val set = sets.getValue(p)

            when(op) {
                1 -> {
                    val y = if (x in set) x else set.firstOrNull() ?: -1
                    res.add(y)
                }
                2 -> {
                    set.remove(x)

                }
            }
        }

        return res.toIntArray()
    }
}

class UnionFind(n: Int) {
    private val parent = IntArray(n) { it }
    private val rank = IntArray(n)

    fun find(x: Int): Int {
        var x = x
        while (parent[x] != x) {
            parent[x] = parent[parent[x]]
            x = parent[x]
        }
        return x
    }

    fun union(x: Int, y: Int) {
        val x = find(x)
        val y = find(y)

        if (x == y) return

        if (rank[x] < rank[y]) {
            parent[x] = y
        } else if (rank[x] > rank[y]) {
            parent[y] = x
        } else {
            parent[y] = x
            rank[x]++
        }
    }
}