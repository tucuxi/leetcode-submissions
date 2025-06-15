class Solution {
    fun isZeroArray(nums: IntArray, queries: Array<IntArray>): Boolean {
        val changes = IntArray(nums.size + 1)
        queries.forEach { (l, r) ->
            changes[l]++
            changes[r + 1]--
        }
        return changes
            .scan(0) { acc, value -> acc + value }
            .drop(1)
            .zip(nums.asIterable())
            .all { (maxOps, num) -> maxOps >= num }        
    }
}