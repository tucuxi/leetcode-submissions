class Solution {
    fun repeatedNTimes(nums: IntArray): Int {
        val v = BooleanArray(10001)
        for (n in nums) {
            if (v[n]) {
                return n
            }
            v[n] = true
        }
        throw IllegalArgumentException("No repeated element")
    }
}