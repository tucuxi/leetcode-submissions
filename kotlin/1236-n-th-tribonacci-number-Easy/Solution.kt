class Solution {
    fun tribonacci(n: Int): Int {
        val t = arrayOf(0, 1, 1)
        repeat(n - t.lastIndex) {
            val next = t.sum()
            t[0] = t[1]
            t[1] = t[2]
            t[2] = next
        }
        return t[minOf(n, t.lastIndex)]
    }
}