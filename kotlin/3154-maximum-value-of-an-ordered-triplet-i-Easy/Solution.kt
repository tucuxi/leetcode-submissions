class Solution {
    fun maximumTripletValue(nums: IntArray): Long {
        var res = 0L
        var ni = 0L
        var d = 0L

        for (num in nums) {
            val n = num.toLong()
            res = maxOf(res, d * n)
            d = maxOf(d, ni - n)
            ni = maxOf(ni, n)
        }

        return res
    }
}