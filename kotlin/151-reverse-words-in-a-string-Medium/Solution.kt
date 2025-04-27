class Solution {
    fun reverseWords(s: String) =
        s.trim().split(" +".toRegex()).reversed().joinToString(" ")
}