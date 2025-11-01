class Solution {
    fun findSmallestInteger(nums: IntArray, value: Int): Int {
        val mp = IntArray(value)
        nums.forEach { n -> mp[n.mod(value)]++ }
        var mex = 0
        while (mp[mex % value] > 0) {
            mp[mex % value]--
            mex++
        }
        return mex
    }
}