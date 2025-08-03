class Solution {
    fun findSubstring(s: String, words: Array<String>): List<Int> {
        val res = mutableListOf<Int>()
        val h = words.groupingBy { it }.eachCount()
        val k = words.first().length

        outer@ for (i in 0 .. s.length - k * words.size) {
            val h2 = mutableMapOf<String, Int>()
            for (w in s.substring(i, i + k * words.size).chunked(k)) {
                h2[w] = h2.getOrDefault(w, 0) + 1
                if (h2.getValue(w) > h.getOrDefault(w, 0)) {
                    continue@outer
                }
            }
            res.add(i)
        }

        return res
    }
}