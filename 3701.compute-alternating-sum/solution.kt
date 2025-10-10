class Solution {
    fun alternatingSum(nums: IntArray): Int {
        var sum = 0
        val iterator = nums.iterator()
        while (iterator.hasNext()) {
            sum += iterator.next()
            if (!iterator.hasNext()) break
            sum -= iterator.next()
        }
        return sum
    }
}