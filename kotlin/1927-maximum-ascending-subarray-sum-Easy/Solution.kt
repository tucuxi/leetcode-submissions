class Solution {
    fun maxAscendingSum(nums: IntArray): Int {
        return nums.runningFold(0 to 0) { (sum, previousNum), num ->
            if (num <= previousNum) {
                num to num
            } else {
                sum + num to num
            }
        }
        .maxBy { it.first }
        .first
    }
}