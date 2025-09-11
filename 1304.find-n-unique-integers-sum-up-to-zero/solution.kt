class Solution {
    fun sumZero(n: Int): IntArray {
        return IntArray(n) { it }
            .apply {
                set(0, -n * (n - 1) / 2)
            }
    }
}