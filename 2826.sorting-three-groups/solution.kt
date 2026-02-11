class Solution {
    fun minimumOperations(nums: List<Int>): Int {
        var dp1 = 0
        var dp2 = 0
        var dp3 = 0

        nums.forEach { 
            when(it) {
                1 -> dp1++
                2 -> dp2 = maxOf(dp1, dp2) + 1
                3 -> dp3 = maxOf(dp1, dp2, dp3) + 1
            }
        }

        return nums.size - maxOf(dp1, dp2, dp3)
    }
}