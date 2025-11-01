class Solution {
    fun maxFrequency(nums: IntArray, k: Int, numOperations: Int): Int {
        var l = 0
        var j = 0
        var r = 0
        val f = mutableMapOf<Int,Int>()

        nums.sort()

        return nums.withIndex().maxOf { (i, x) ->
            while (r < nums.size && nums[r] <= x + k) {
                nums[r].let { f[it] = f.getOrDefault(it, 0) + 1 }
                r++
            }
            while (l < nums.size && nums[l] < x - k) {
                nums[l].let { f[it] = f.getValue(it) - 1 }
                l++
            }
            while (x - nums[j] > 2 * k) {
                j++
            }
            maxOf(
                minOf(f.getValue(x) + numOperations, r - l),
                minOf(i - j + 1, numOperations)
            )
        }
    }
}