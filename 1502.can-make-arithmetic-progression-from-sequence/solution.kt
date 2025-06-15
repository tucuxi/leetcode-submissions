class Solution {
    fun canMakeArithmeticProgression(arr: IntArray): Boolean {
        val sorted = arr.sorted()
        val delta = sorted.zip(sorted.drop(1)).map { it.first - it.second }
        return delta.all { it == delta.first() }
    }
}