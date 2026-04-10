class Solution {
    fun checkStrings(s1: String, s2: String): Boolean {
        return checkFrom(s1, s2, 0) && checkFrom(s1, s2, 1)
    }

    private fun checkFrom(s1: String, s2: String, start: Int): Boolean {
        val h = IntArray(26)
        for (i in start..s1.lastIndex step 2) {
            h[s1[i] - 'a']++
            h[s2[i] - 'a']--
        }
        return h.all { it == 0 }
    }
}