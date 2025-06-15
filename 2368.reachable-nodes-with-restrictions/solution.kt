class Solution {
    fun reachableNodes(n: Int, edges: Array<IntArray>, restricted: IntArray): Int {
        val restrictedNodes = restricted.toSet()

        val nodes = edges
            .filterNot { (a, b) -> a in restrictedNodes || b in restrictedNodes }
            .flatMap { (a, b) -> listOf(a to b, b to a) }
            .groupBy { (k, _) -> k }
            .mapValues { (_, v) -> v.map { it.second } }

        fun dfs(nodeIndex: Int, parentIndex: Int): Int {
            return 1 + nodes.getOrDefault(nodeIndex, emptyList())
                .filterNot { childIndex -> childIndex == parentIndex }
                .sumOf { childIndex -> dfs(childIndex, nodeIndex) }
        }

        return dfs(0, -1)
    }
}