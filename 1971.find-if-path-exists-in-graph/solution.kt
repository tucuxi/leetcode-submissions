class Solution {
    fun validPath(n: Int, edges: Array<IntArray>, start: Int, end: Int): Boolean {
        val visited = BooleanArray(n)
        val neighbors = Array(n) { mutableSetOf<Int>() }
        
        for ((a, b) in edges) {
            neighbors[a].add(b)
            neighbors[b].add(a)
        }
        
        fun dfs(node: Int): Boolean {
            if (node == end) {
                return true
            }
            if (visited[node]) {
                return false
            }
            visited[node] = true
            for (neighbor in neighbors[node]) {
                if (dfs(neighbor)) {
                    return true
                }
            }
            return false
        }
        
        return dfs(start)
    }
}