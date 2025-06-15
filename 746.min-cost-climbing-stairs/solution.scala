object Solution {
    def minCostClimbingStairs(cost: Array[Int]): Int = {
        val t = Array.fill[Int](cost.length + 1)(0)
        for (i <- 2 to cost.length)
            t(i) = math.min(t(i - 1) + cost(i - 1), t(i - 2) + cost(i - 2))
        t(cost.length)
    }
}