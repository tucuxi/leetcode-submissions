class Solution {
    fun longestOnes(nums: IntArray, k: Int): Int {
        var i = 0
        var r = k

        nums.forEach { num ->
            if (num == 0) {
                r--
            }
            if (r < 0) {
                if (nums[i] == 0) {
                    r++
                }
                i++
            }
        }
        return nums.size - i
    }
}
