class Solution {
    fun minimumDistance(nums: IntArray): Int {
        var res = Int.MAX_VALUE
        val first = mutableMapOf<Int, Int>()
        val second = mutableMapOf<Int, Int>()

        nums.forEachIndexed { i, n ->
            if (n in second) {
                res = minOf(res, 2 * (i - first.getValue(n)))
                first[n] = second.getValue(n)
                second[n] = i
            } else if (n in first) {
                second[n] = i
            } else {
                first[n] = i
            }
        }

        return if (res == Int.MAX_VALUE) -1 else res
    }
}