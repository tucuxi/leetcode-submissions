private const val MOD = 1_000_000_007

class Solution {
    fun minAbsoluteSumDiff(nums1: IntArray, nums2: IntArray): Int {
        val s = TreeSet<Int>(nums1.asList())
        val pairs  = nums1.zip(nums2)
        val sum = pairs.sumOf { (a, b) -> abs(a - b).toLong() }

        val res = pairs.minOf { (a, b) ->
            listOfNotNull(
                s.floor(b),
                s.ceiling(b)
            )
            .map { sum - abs(a - b).toLong() + abs(it - b).toLong() }
            .minOrNull()
            ?: sum
        }

        return (res % MOD).toInt()
    }
}