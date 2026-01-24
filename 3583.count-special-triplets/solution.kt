private const val MOD = 1_000_000_007

class Solution {
    fun specialTriplets(nums: IntArray): Int {
        val prefix = mutableMapOf<Int, Int>()
        val suffix = nums.asIterable().groupingBy { it }.eachCount().toMutableMap()
        var res = 0L

        nums.forEach { num ->
            suffix[num] = suffix.getValue(num) - 1
            res += prefix.getOrDefault(2 * num, 0).toLong() * suffix.getOrDefault(2 * num, 0).toLong()
            res %= MOD
            prefix[num] = prefix.getOrPut(num, { 0 }) + 1
        }

        return res.toInt()
    }
}