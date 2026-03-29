class Solution {
    fun canBeEqual(s1: String, s2: String): Boolean {
        val c1 = s1.toCharArray()
        if (c1[0] != s2[0]) {
            c1[0] = c1[2].also { c1[2] = c1[0] }
        }
        if (c1[1] != s2[1]) {
            c1[1] = c1[3].also { c1[3] = c1[1] }
        }
        return String(c1) == s2
    }
}