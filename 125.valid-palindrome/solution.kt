class Solution {
    fun isPalindrome(s: String): Boolean {
        val normalized = s.mapNotNull { ch -> when {
            ch.isLetter() -> ch.toLowerCase()
            ch.isDigit() -> ch
            else -> null
        } }.joinToString(separator = "")
        return isPalindromeRec(normalized)
    }
    
    private fun isPalindromeRec(s: String): Boolean =
        when {
            s.length <= 1 -> true
            s.first() != s.last() -> false
            else -> isPalindromeRec(s.substring(1, s.length - 1))
        }
}