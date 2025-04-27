class Solution {
    fun lengthOfLongestSubstring(s: String): Int {
        val lastOccurrence = HashMap<Char, Int>()
        var longest = 0
        var run = 0
        s.forEachIndexed { i, c ->
            lastOccurrence.get(c)?.let { j ->
                if (i - j <= run) {
                    run = i - j - 1
                }
            }
            lastOccurrence.put(c, i)
            longest = maxOf(++run, longest)
        }
        return longest
    }
}