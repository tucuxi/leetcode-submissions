class Solution {
    fun maxKDivisibleComponents(n: Int, edges: Array<IntArray>, values: IntArray, k: Int): Int {
        
        val adj = Array(n) { mutableListOf<Int>() }

        edges.forEach { (u, v) ->
            adj[u].add(v)
            adj[v].add(u)
        }
        
        var components = 0

        fun dfs(currentNode: Int, parentNode: Int): Int {
            var sum = 0

            adj[currentNode].forEach { neighborNode ->
                if (neighborNode != parentNode) {
                    sum = (sum + dfs(neighborNode, currentNode)) % k
                }
            }

            sum = (sum + values[currentNode]) % k

            if (sum == 0) {
                components++
            }

            return sum
        }

        dfs(0, -1)
        return components
    }
}