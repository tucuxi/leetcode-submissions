class Solution {
    fun maximumSum(nums: IntArray): Int {
        var res = -1
        val h = IntArray(82)

        nums.forEach { n ->
            val ds = digitsSum(n)
            if (h[ds] > 0) {
                res = maxOf(res, h[ds] + n)
            }
            if (n > h[ds]) {
                h[ds] = n
            }
        }
        return res
    }

    private fun digitsSum(num: Int): Int {
        var res = 0
        var n = num

        while (n > 0) {
            res += n % 10
            n /= 10
        }
        return res
    }
}