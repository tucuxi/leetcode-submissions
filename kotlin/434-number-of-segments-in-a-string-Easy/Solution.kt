class Solution {
    fun countSegments(s: String): Int {
        var count = 0
        var inside = false
        s.forEach { ch ->
            if (ch == ' ') {
                inside = false
            } else {
                if (!inside) count++
                inside = true
            }
        }
        return count
    }
}