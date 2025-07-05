class Solution {
    fun findLHS(nums: IntArray): Int {
        val h = nums.asSequence().groupingBy { it }.eachCount()
        var res = 0
        h.forEach { (min, minCount) ->
            h[min + 1]?.let { maxCount ->
                res = maxOf(res, minCount + maxCount)
            }
        }
        return res
    }
}