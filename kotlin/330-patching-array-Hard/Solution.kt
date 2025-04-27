class Solution {
    fun minPatches(nums: IntArray, n: Int): Int {
        var i = 0
        var res = 0
        var expect = 1L
        while (expect <= n) {
            if (i >= nums.size || nums[i] > expect) {
                res++
                expect += expect
            } else {
                expect += nums[i++]
            }
        }
        return res
    }
}