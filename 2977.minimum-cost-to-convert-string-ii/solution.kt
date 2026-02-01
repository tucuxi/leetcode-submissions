const val INF = Long.MAX_VALUE / 2

class Solution {
    fun minimumCost(source: String, target: String, original: Array<String>, changed: Array<String>, cost: IntArray): Long {
        val strToId = mutableMapOf<String, Int>()
        var idCounter = 0

        fun getId(s: String): Int = strToId.getOrPut(s) { idCounter++ }

        val distinctLengths = mutableSetOf<Int>()

        for (s in original) {
            getId(s)
            distinctLengths.add(s.length)
        }
        for (s in changed) {
            getId(s)
        }

        val numNodes = idCounter
        val dist = Array(numNodes) { LongArray(numNodes) { INF } }

        for (i in 0 until numNodes) {
            dist[i][i] = 0
        }
        for (i in original.indices) {
            val u = getId(original[i])
            val v = getId(changed[i])
            dist[u][v] = min(dist[u][v], cost[i].toLong())
        }
        for (k in 0 until numNodes) {
            for (i in 0 until numNodes) {
                if (dist[i][k] == INF) continue
                for (j in 0 until numNodes) {
                    if (dist[k][j] == INF) continue
                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
                }
            }
        }

        val n = source.length
        val dp = LongArray(n + 1) { INF }
        dp[0] = 0

        for (i in 0 until n) {
            if (dp[i] == INF) continue

            if (source[i] == target[i]) {
                dp[i + 1] = min(dp[i + 1], dp[i])
            }
            for (len in distinctLengths) {
                if (i + len > n) continue
                val subSrc = source.substring(i, i + len)
                val subTgt = target.substring(i, i + len)
                val u = strToId[subSrc]
                val v = strToId[subTgt]
                if (u != null && v != null && dist[u][v] < INF) {
                    dp[i + len] = min(dp[i + len], dp[i] + dist[u][v])
                }
            }
        }
        return if (dp[n] >= INF) -1 else dp[n]
    }
}