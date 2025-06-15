class Solution {
    fun firstUniqChar(s: String): Int {
        val freq = IntArray(26)
        for (ch in s) freq[ch - 'a']++
        for (i in 0 .. s.length - 1) {
            if (freq[s[i] - 'a'] == 1) return i 
        }
        return -1
    }
}