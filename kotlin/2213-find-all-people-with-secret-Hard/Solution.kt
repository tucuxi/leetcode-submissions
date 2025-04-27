class Solution {
    fun findAllPeople(n: Int, meetings: Array<IntArray>, firstPerson: Int): List<Int> {
        var chrono = meetings.sortedBy { it[2] }
        val uf = DisjointSet(n)
        uf.union(0, firstPerson)
        
        var i = 0
        while (i < chrono.size) {
            var j = i
            while (j < chrono.size && chrono[j][2] == chrono[i][2]) {
                uf.union(chrono[j][0], chrono[j][1])
                j++
            }
            while (i < j) {
                if (!uf.connected(chrono[i][0], 0)) {
                    uf.reset(chrono[i][0])
                    uf.reset(chrono[i][1])
                }
                i++
            }
        }
        
        return (0 until n).filter { uf.connected(it, 0) }
    }
}

class DisjointSet(n : Int) {
    private val parent = IntArray(n) { it }
    private val rank = IntArray(n)
    
    override fun toString() = "DisjointSet(parent=[${parent.joinToString()}], rank=[${rank.joinToString()}])"

    fun union(a: Int, b: Int) {
        val ap = find(a)
        val bp = find(b)
        if (ap == bp) {
            return
        }
        if (rank[ap] < rank[bp]) {
            parent[ap] = bp
        } else if (rank[ap] > rank[bp]) {
            parent[bp] = ap
        } else {
            parent[bp] = ap
            rank[ap]++
        }
    }
    
    fun find(a: Int): Int {
        if  (a != parent[a]) {
            parent[a] = find(parent[a])
        }
        return parent[a]
    }
    
    fun connected(a: Int, b: Int) = find(a) == find(b)
    
    fun reset(a: Int) {
        parent[a] = a
        rank[a] = 0
    }
}