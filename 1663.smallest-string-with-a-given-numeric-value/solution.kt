class Solution {
    fun getSmallestString(n: Int, k: Int): String {
        val res = CharArray(n)
        var rest = k

        for (i in n - 1 downTo 0) {
            val current = minOf(rest - i, 26)
            res[i] = 'a' + current - 1
            rest -= current
        }
        return String(res)
    }
}