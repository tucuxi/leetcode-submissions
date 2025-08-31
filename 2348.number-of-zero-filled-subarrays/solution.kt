class Solution {
    fun zeroFilledSubarray(nums: IntArray): Long {
        var res = 0L
        var run = 0
        nums.forEach { num ->
            if (num == 0) {
                run++
                res += run
            } else {
                run = 0
            }
        }
        return res
    }
}