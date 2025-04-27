class Solution {
    fun countQuadruplets(nums: IntArray): Int {
        var res = 0
        for (a in 0 until nums.size) {
            for (b in a+1 until nums.size) {
                for (c in b+1 until nums.size) {
                    val sum = nums[a] + nums[b] + nums[c]
                    for (d in c+1 until nums.size) {
                        if (sum == nums[d]) {
                            res++
                        }
                    }
                }
            }
        }
        return res
    }
}