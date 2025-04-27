class Solution {
    fun backspaceCompare(s: String, t: String): Boolean {
        return norm(s) == norm(t)
    }
    
    private fun norm(s: String): String {
        val b = StringBuilder()
        var skip = 0
        for (i in s.lastIndex downTo 0) {
            when {
                s[i] == '#' -> skip++
                skip > 0 -> skip--
                else -> b.append(s[i])
            }
        }
        return b.toString()
    }
}