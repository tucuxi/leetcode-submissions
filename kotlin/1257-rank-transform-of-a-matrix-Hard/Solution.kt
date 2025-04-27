class Solution {
    fun matrixRankTransform(matrix: Array<IntArray>): Array<IntArray> {
        val m = matrix.size
        val n = matrix[0].size
        
        fun makeElementMap(): Map<Int, List<Pair<Int, Int>>> {
            val res = mutableMapOf<Int, MutableList<Pair<Int, Int>>>()
            for (i in 0 until m) {
                for (j in 0 until n) {
                    val element = matrix[i][j]
                    val list = res[element]
                    if (list != null) {
                        list.add(Pair(i, j))
                    } else {
                        res[element] = mutableListOf(Pair(i, j))
                    }
                }
            }
            return res
        }
        
        val rank = IntArray(m + n)
        val elementMap = makeElementMap()
        for (key in elementMap.keys.sorted()) {
            val elements = elementMap[key]!!
            val dsu = DSU(elements.map { (i, _) -> i } + elements.map { (_, j) -> j + m } )
            elements.forEach { (i, j) -> dsu.union(i, j + m) }
            for (group in dsu.groups().values) {
                val mx = 1 + group.map { rank[it] }.max()!!
                group.forEach { rank[it] = mx }
            }
            elements.forEach { (i, j) -> matrix[i][j] = rank[i] }
        }
        return matrix
    }
}

private class DSU(elements: List<Int>) {
    val parent = elements.map { it to it }.toMap().toMutableMap()
    
    fun find(element: Int): Int {
        if (parent[element] != element) {
            parent.put(element, find(parent[element]!!))
        }
        return parent[element]!!
    }
    
    fun union(element1: Int, element2: Int) {
        parent[find(element1)] = find(element2)
    }
    
    fun groups(): Map<Int, List<Int>> {
        val res = mutableMapOf<Int, MutableList<Int>>()
        for (element in parent.keys) {
            val group = find(element)
            val list = res[group]
            if (list != null) {
                list.add(element)
            } else {
                res[group] = mutableListOf(element)
            }
        }
        return res
    }
}