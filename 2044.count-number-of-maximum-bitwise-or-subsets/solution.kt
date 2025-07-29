class Solution {
    fun countMaxOrSubsets(nums: IntArray): Int {
        
        fun dfs(i: Int, pre: Int): Int =
            if (i < nums.size) {
                dfs(i + 1, pre) + dfs(i + 1, pre and nums[i].inv())
            } else if (pre == 0) {
                1
            } else {
                0
            }

        val maxOr = nums.fold(0) { acc, num -> acc or num }
        return dfs(0, maxOr)
    }
}