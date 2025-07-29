class Solution {
    fun maxSum(nums: IntArray): Int {
        val (nonNeg, neg) = nums.partition { it >= 0 }
        return if (nonNeg.isNotEmpty()) {
            nonNeg.distinct().sum()
         } else {
            neg.max()
         } 
    }
}