const val VOWELS = "AEIOUaeiou"

class Solution {
    fun reverseVowels(s: String): String {
        val sb = StringBuilder(s)
        val n = s.length
        var i = 0
        var j = n - 1
        while (i < j) {
            while (i < j && !VOWELS.contains(sb[i])) {
                i++
            }
            while (j > i && !VOWELS.contains(sb[j])) {
                j--
            }
            if (i == j) {
                break
            }
            sb[i] = sb[j].also { sb[j] = sb[i] }
            i++
            j--
        }
        return sb.toString()
    }
}