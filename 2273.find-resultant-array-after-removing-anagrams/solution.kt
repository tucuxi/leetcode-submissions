class Solution {
    fun removeAnagrams(words: Array<String>): List<String> {
        val res = mutableListOf(words[0])
        for (i in 1 until words.size) {
            if (!isAnagram(words[i - 1], words[i])) {
                res.add(words[i])
            }
        }
        return res
    }

    private fun isAnagram(s1: String, s2: String): Boolean {
        if (s1.length != s2.length) {
            return false
        }
        val count = IntArray(26)
        for (i in s1.indices) {
            count[s1[i] - 'a']++
            count[s2[i] - 'a']--
        }
        for (c in count) {
            if (c != 0) {
                return false
            }
        }
        return true
    }
}