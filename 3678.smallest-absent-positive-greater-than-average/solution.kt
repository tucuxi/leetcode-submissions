class Solution {
    fun smallestAbsent(nums: IntArray): Int {
        val v = BooleanArray(102)
        val sum = nums.fold(0) { acc, num ->
            if (num > 0) {
                v[num] = true
            }
            acc + num
        }
        val a = maxOf(0, sum / nums.size) + 1
        return (a..v.lastIndex).first { !v[it] }
    }
}