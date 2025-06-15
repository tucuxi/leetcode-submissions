class Solution {
    fun calcEquation(equations: List<List<String>>, values: DoubleArray, queries: List<List<String>>): DoubleArray {
        val graph = mutableMapOf<String, MutableList<Pair<String, Double>>>().apply {
            equations.forEachIndexed { index, (a, b) ->
                if (!contains(a)) {
                    put(a, mutableListOf(b to values[index]))
                } else {
                    get(a)!!.add(b to values[index])
                }
                if (!contains(b)) {
                    put(b, mutableListOf(a to 1 / values[index]))
                } else {
                    get(b)!!.add(a to 1 / values[index])
                }
            }
        }
        
        fun dfs(a: String, b: String, visited: MutableSet<String>): Double {
            if (a in visited || a !in graph) {
                return -1.0
            }
            visited.add(a)
            if (a == b) {
                return 1.0
            }
            graph.get(a)!!.forEach { (c, v) ->
                val x = dfs(c, b, visited)
                if (x > 0) {
                    return x * v
                }
            }
            return -1.0
        }
        
        return queries.map { (p, q) -> dfs(p, q, mutableSetOf<String>()) }.toDoubleArray()
    }
}