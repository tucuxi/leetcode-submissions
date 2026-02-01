class Solution {
    fun pathExistenceQueries(n: Int, nums: IntArray, maxDiff: Int, queries: Array<IntArray>): BooleanArray {
        var currentGroup = 0
        val group = IntArray(n).also { it[0] = currentGroup }
        
        for (i in 1 until n) {
            if (nums[i] - nums[i - 1] > maxDiff) {
                currentGroup++
            }
            group[i] = currentGroup
        }

        return queries.map { (u, v) -> group[u] == group[v] }.toBooleanArray()
    }
}