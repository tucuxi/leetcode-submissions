class Solution {
    fun maxProfit(n: Int, present: IntArray, future: IntArray, hierarchy: Array<IntArray>, budget: Int): Int {
        val g = Array(n) { mutableListOf<Int>() }
        hierarchy.forEach { (u, v) -> g[u - 1].add(v - 1) }
        
        fun dfs(u: Int): Result {
            val cost = present[u]
            val dCost = present[u] / 2
            val dp0 = IntArray(budget + 1)
            val dp1 = IntArray(budget + 1)
            val subProfit0 = IntArray(budget + 1)
            val subProfit1 = IntArray(budget + 1)
            var uSize = cost

            g[u].forEach { v ->
                val childResult = dfs(v)
                uSize += childResult.size
                for (i in budget downTo 0) {
                    for (sub in 0 .. minOf(childResult.size, i)) {
                        if (i - sub >= 0) {
                            subProfit0[i] = maxOf(
                                subProfit0[i],
                                subProfit0[i - sub] + childResult.dp0[sub]
                            )
                            subProfit1[i] = maxOf(
                                subProfit1[i],
                                subProfit1[i - sub] + childResult.dp1[sub]
                            )
                        }
                    }
                }
            }

            for (i in 0 .. budget) {
                dp0[i] = subProfit0[i]
                dp1[i] = subProfit0[i]
                if (i >= dCost) {
                    dp1[i] = maxOf(
                        subProfit0[i],
                        subProfit1[i - dCost] + future[u] - dCost
                    )
                }
                if (i >= cost) {
                    dp0[i] = maxOf(
                        subProfit0[i],
                        subProfit1[i - cost] + future[u] - cost
                    )
                }
            }
            return Result(dp0, dp1, uSize)
        }
    
        return dfs(0).dp0[budget]
    }

    class Result(val dp0: IntArray, val dp1: IntArray, val size: Int)
}