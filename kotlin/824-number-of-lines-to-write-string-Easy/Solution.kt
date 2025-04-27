class Solution {
    fun numberOfLines(widths: IntArray, s: String): IntArray {
        var line = 1
        var pixels = 0
        for (ch in s) {
            val width = widths[ch - 'a']
            if (pixels + width > 100) {
                line++
                pixels = width
            } else {
                pixels += width
            }
        }
        return intArrayOf(line, pixels)
    }
}