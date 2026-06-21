class Solution {
    fun earliestFinishTime(landStartTime: IntArray, landDuration: IntArray, waterStartTime: IntArray, waterDuration: IntArray): Int {
        return minOf(
            optimize(landStartTime, landDuration, waterStartTime, waterDuration),
            optimize(waterStartTime, waterDuration, landStartTime, landDuration)
        )
    }

    fun optimize(s1: IntArray, d1: IntArray, s2: IntArray, d2: IntArray): Int {
        val t1 = (s1 zip d1).minOf { (s, d) -> s + d }
        return (s2 zip d2).minOf { (s, d) -> maxOf(t1, s) + d }
    }
}