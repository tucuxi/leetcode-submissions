class Solution {
    fun maxSum(nums: IntArray): Int {
        return nums.partition { it >= 0 }.let { (p, n) ->
            if (p.isNotEmpty()) {
                p.toSet().sum()
            } else {
                n.max()
            }
        }
    }
}