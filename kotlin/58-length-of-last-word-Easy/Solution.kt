class Solution {
    fun lengthOfLastWord(s: String): Int {
        val t = s.trim(' ')
        val lastSpacePos = t.lastIndexOf(' ')
        return if (lastSpacePos == -1) t.length else t.length - lastSpacePos - 1
    }
}