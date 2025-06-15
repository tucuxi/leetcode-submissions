class Solution {
    fun modifyString(s: String): String {
        val res = StringBuilder(s)
        for (i in res.indices) {
            if (res[i] == '?') {
                res[i] = otherChar(
                    res.getOrElse(i - 1) { 'a' },
                    res.getOrElse(i + 1) { 'a' }
                )       
            }
        }
        return res.toString()
    }
}

fun otherChar(ch1: Char, ch2: Char): Char {
    if (ch1 != 'a' && ch2 != 'a') return 'a'
    if (ch1 != 'b' && ch2 != 'b') return 'b'
    return 'c'
}