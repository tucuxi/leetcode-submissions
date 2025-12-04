class Solution {
    fun minMoves(nums: IntArray): Int {
        val (sum, maxNum) = nums.fold(0 to 0) { (sum, maxNum), num ->
            sum + num to maxOf(maxNum, num)
        }
        return nums.size * maxNum - sum
    }
}