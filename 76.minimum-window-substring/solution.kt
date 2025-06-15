class Solution {
    fun minWindow(s: String, t: String): String {
        val restMap = t.groupingBy{ it }.eachCount().toMutableMap()
        var restCount = t.length
        var left = 0
        var minLeft = 0
        var minLen = s.length + 1
        for (right in 0 until s.length) {
            val charRight = s[right]
            restMap.get(charRight)?.let {
                restMap.put(charRight, it - 1)
                if (it > 0) {
                    restCount--
                }
            }
            while (restCount == 0) {
                val len = right - left + 1
                if (len < minLen) {
                    minLen = len
                    minLeft = left
                } 
                val charLeft = s[left]
                restMap.get(charLeft)?.let {
                    restMap.put(charLeft, it + 1)
                    if (it >= 0) {
                        restCount++
                    }
                }
                left++
            }
        }
        return if (minLen > s.length) "" else s.substring(minLeft, minLeft + minLen)
    }
}