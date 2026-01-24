private const val MOD = 1_000_000_007

class Solution {
    fun xorAfterQueries(nums: IntArray, queries: Array<IntArray>): Int {
        queries.forEach { (l, r, k, v) ->
            for (i in l..r step k) {
                nums[i] = (nums[i].toLong() * v % MOD).toInt()
            }
        }
        return nums.fold(0) { acc, num -> acc xor num }
    }
}