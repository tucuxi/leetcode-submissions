class Solution {
    fun subsetXORSum(nums: IntArray): Int {
        fun rec(acc: Int, index: Int): Int =
            if (index == nums.size) {
                acc
            } else {
                rec(acc xor nums[index], index + 1) + rec(acc, index + 1)
        }
        return rec(0, 0)
    }

}