class Solution {
    fun maxPower(s: String): Int {
        if (s.length == 0) return 0
        var currentLen = 1
        var prevChar = s[0]
        var maxLen = 1    
        for (i in 1 until s.length) {
            val currentChar = s[i]
            if (currentChar == prevChar) {
                if (++currentLen > maxLen) {
                    maxLen = currentLen
                }
            } else {
                prevChar = currentChar
                currentLen = 1
            }
        }
        return maxLen
    }
}