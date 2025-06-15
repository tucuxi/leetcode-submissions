import java.util.PriorityQueue

class Solution {
    fun smallestStringWithSwaps(s: String, pairs: List<List<Int>>): String {
        val dsu = DSU(s.length).apply {
            pairs.forEach { (a, b) -> union(a, b) }
        }
        val map: MutableMap<Int, PriorityQueue<Char>> = mutableMapOf()
        s.forEachIndexed { i, c ->
            val component = dsu.find(i)
            if (component !in map) {
                map[component] = PriorityQueue()
            }
            map[component]!!.offer(c)
        }
        val res = StringBuilder(s.length)
        s.indices.forEach { i ->
            val component = dsu.find(i)
            val chars = map[component]!!
            val c = chars.poll()
            res.append(c)
        }
        return res.toString()
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