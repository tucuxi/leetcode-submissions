class Solution {
    fun repeatedSubstringPattern(s: String): Boolean {
        return (1 .. s.length / 2).any { n->
            s.length % n == 0 && s == s.take(n).repeat(s.length / n)
        }
    }
}