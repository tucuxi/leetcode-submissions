class Solution {
    fun getLongestSubsequence(words: Array<String>, groups: IntArray): List<String> {
        return words.filterIndexed { i, _ -> i == 0 || groups[i] != groups[i - 1] }
    }
}