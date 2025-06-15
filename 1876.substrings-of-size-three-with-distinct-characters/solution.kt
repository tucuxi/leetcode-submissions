class Solution {
    fun countGoodSubstrings(s: String): Int =
        s.windowed(3).count { t -> t[0] != t[1] && t[0] != t[2] && t[1] != t[2] }
}