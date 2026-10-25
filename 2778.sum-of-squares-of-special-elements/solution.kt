class Solution {
    fun sumOfSquares(nums: IntArray): Int {
        return nums.withIndex().sumOf { (i, x) ->
            if (nums.size % (i + 1) == 0)
                x * x
            else
                0
        }
    }
}