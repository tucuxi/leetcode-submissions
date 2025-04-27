class Solution {
    fun getLongestSubsequence(words: Array<String>, groups: IntArray): List<String> {
        val res = mutableListOf(words[0])
        for (i in 1 until groups.size) {
            if (groups[i] != groups[i-1]) {
                res.add(words[i])
            }
        }
        return res
    }
}