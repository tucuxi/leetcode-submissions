class Solution {
    fun mergeCharacters(s: String, k: Int): String {
        val res = StringBuilder()
        val occurrence = mutableMapOf<Char, Int>()
        for (ch in s) {
            val i = occurrence[ch]
            val j = res.length
            if (i == null || j - i > k) {
                res.append(ch)
                occurrence[ch] = j
            }
        }
        return res.toString()
    }
}