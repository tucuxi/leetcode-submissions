class Solution {
    fun minBitwiseArray(nums: List<Int>): IntArray {
        return IntArray(nums.size) { i ->
            var b = 1
            var res = -1
            while (nums[i] and b != 0) {
                res = nums[i] - b
                b = b shl 1
            }
            res
        }
    }
}