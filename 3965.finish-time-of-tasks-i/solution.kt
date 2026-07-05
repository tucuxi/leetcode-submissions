class Solution {

    fun finishTime(n: Int, edges: Array<IntArray>, baseTime: IntArray): Long {
        val children = Array(n) { mutableListOf<Int>() }
        edges.forEach { (u, v) -> children[u] += v }

        fun dfs(task: Int): Long {
            return if (children[task].isEmpty()) {
                baseTime[task].toLong()
            } else {
                val finishTimes = children[task].map { child -> dfs(child) }
                val earliest = finishTimes.min()
                val latest = finishTimes.max()
                val ownDuration = latest - earliest + baseTime[task]
                latest + ownDuration
            }
        }

        return dfs(0)
    }
}