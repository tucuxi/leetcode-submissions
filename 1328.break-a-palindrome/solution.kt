class Solution {
    fun breakPalindrome(palindrome: String): String {
        if (palindrome.length <= 1) return ""
        val res = StringBuilder(palindrome)
        val n = palindrome.length

        for (i in 0 until n / 2) {
            if (res[i] != 'a') 
            {
                res[i] = 'a'
                return res.toString()
            }
        }
        res[n - 1] = 'b'
        return res.toString()
    }
}