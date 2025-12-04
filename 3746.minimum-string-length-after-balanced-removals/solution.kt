class Solution {
    fun minLengthAfterRemovals(s: String): Int {
        val d = s.fold(0) { acc, ch -> acc + if (ch == 'a') 1 else -1 }
        return abs(d)
    }
}