class Solution {
    fun removePalindromeSub(s: String): Int {
        return if ((0 .. s.length / 2).all { s[it] == s[s.lastIndex - it] }) 1 else 2
    }
}