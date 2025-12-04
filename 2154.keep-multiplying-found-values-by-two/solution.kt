class Solution {
    fun findFinalValue(nums: IntArray, original: Int): Int {
        val v = BooleanArray(1001)
        nums.forEach { v[it] = true }
        var n = original
        while (n < v.size && v[n]) {
            n += n
        }
        return n
    }
}