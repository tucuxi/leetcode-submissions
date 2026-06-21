class Solution {
    fun earliestFinishTime(landStartTime: IntArray, landDuration: IntArray, waterStartTime: IntArray, waterDuration: IntArray): Int {
        return minOf(
            optimize(landStartTime, landDuration, waterStartTime, waterDuration),
            optimize(waterStartTime, waterDuration, landStartTime, landDuration)
        )
    }

    fun optimize(s1: IntArray, d1: IntArray, s2: IntArray, d2: IntArray): Int {
        val t1 = (0 until s1.size).minOf { s1[it] + d1[it] }
        return (0 until s2.size).minOf { maxOf(t1, s2[it]) + d2[it] }
    }
}