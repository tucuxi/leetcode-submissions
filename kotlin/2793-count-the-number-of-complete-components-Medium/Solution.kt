class Solution {
    fun countCompleteComponents(n: Int, edges: Array<IntArray>): Int {
        val indegree = edges
            .flatMap { (a, b) -> listOf(a to b, b to a) }
            .groupBy { it.second }
            .mapValues { it.value.size }

        val uf = UnionFind(n)
        edges.forEach { (a, b) -> uf.union(a, b) }

        return uf.components().count { a ->
            val members = uf.members(a) 
            members.all { x -> indegree.getOrDefault(x, 0) == members.size - 1 }
        }
    }
}

class UnionFind(n: Int) {
    private val parent = IntArray(n) { it }
    private val count = IntArray(n) { 1 }

    fun union(a: Int, b: Int) {
        val x = find(a)
        val y = find(b)
        if (x != y) {
            if (count[x] < count[y]) {
                parent[x] = y
                count[y] += count[x]
            } else {
                parent[y] = x
                count[x] += count[y]
            }
        }
    }

    fun find(a: Int): Int {
        if (parent[a] != a) {
            parent[a] = find(parent[a])
        }
        return parent[a]
    }

    fun components(): Iterable<Int> {
        return parent.toSet()
    }

    fun members(a: Int): List<Int> {
        return parent.withIndex()
            .filter { (_, x) -> x == a }
            .map { (index, _) -> index }
    }
}