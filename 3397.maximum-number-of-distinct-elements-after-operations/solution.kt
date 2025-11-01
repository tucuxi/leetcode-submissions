class Solution {
    fun maxDistinctElements(nums: IntArray, k: Int): Int {
        var p = Int.MIN_VALUE
        var res = 0

        for (num in nums.sorted()) {
            val q = maxOf(p, num - k)
            if (q <= num + k) {
                res++
                p = q + 1
            }
        }
        return res
    }
}