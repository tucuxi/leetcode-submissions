class Solution {
    fun pathExistenceQueries(n: Int, nums: IntArray, maxDiff: Int, queries: Array<IntArray>): BooleanArray {
        val group = IntArray(n)
        var currentGroup = 0

        for (i in 1 until n) {
            if (nums[i] - nums[i - 1] > maxDiff) {
                currentGroup++
            }
            group[i] = currentGroup
        }

        return queries.map { (u, v) -> group[u] == group[v] }.toBooleanArray()
    }
}