class Solution {
    fun largestSumAfterKNegations(nums: IntArray, k: Int): Int {
        repeat(k) {
            var min  = nums.first()
            var minIndex = 0
            nums.indices.forEach {
                if (nums[it] < min) {
                    min = nums[it]
                    minIndex = it
                }
            }
            nums[minIndex] = -nums[minIndex]
        }
        return nums.sum()
    }
}