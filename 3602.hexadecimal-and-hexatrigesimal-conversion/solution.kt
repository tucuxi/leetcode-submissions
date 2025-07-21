class Solution {
    fun concatHex36(n: Int): String {
        val s = (n * n).toString(16) + (n * n * n).toString(36)
        return s.uppercase()
    }
}