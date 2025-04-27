class Solution {
    fun xorOperation(n: Int, start: Int): Int {
        val nums = IntArray(n) { start + 2 * it }
        return nums.reduce { a, b -> a.xor(b) }
    }
}