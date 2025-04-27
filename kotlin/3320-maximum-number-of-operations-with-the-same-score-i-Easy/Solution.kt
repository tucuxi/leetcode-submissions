class Solution {
    fun maxOperations(nums: IntArray): Int {
        val sum = nums[0] + nums[1]
        var res = 1
        var i = 2
        
        while (i < nums.lastIndex && nums[i] + nums[i+1] == sum) {
            res++
            i += 2
        }
        return res
    }
}