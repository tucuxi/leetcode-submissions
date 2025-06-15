class Solution {
    fun countPairs(nums: IntArray, k: Int): Int {
        var count = 0
        nums.forEachIndexed { i, n ->
            for (j in i+1 .. nums.lastIndex) {
                if (n == nums[j] && i * j % k == 0) {
                    count++
                }
            }
        }
        return count
    }
}