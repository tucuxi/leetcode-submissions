class Solution {
    fun countOppositeParity(nums: IntArray): IntArray {
        val res = IntArray(nums.size)
        var even = 0
        var odd = 0

        for (i in nums.lastIndex downTo 0) {
            if (nums[i] % 2 == 0) {
                res[i] = odd
                even++
            } else {
                res[i] = even
                odd++
            }
        }

        return res
    }
}