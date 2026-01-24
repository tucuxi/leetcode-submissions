class Solution {
    fun reversePrefix(s: String, k: Int): String {
        val pre = s.substring(0, k)
        return pre.reversed() + s.substring(k)
    }
}