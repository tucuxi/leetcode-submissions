class Solution {
    fun shiftingLetters(s: String, shifts: IntArray): String {
        var n = 0
        val t = StringBuilder(s)
        for (i in t.indices.reversed()) {
            n = (n + shifts[i]) % 26
            t[i] = shift(t[i], n)
        }
        return t.toString()
    }
}

inline fun shift(ch: Char, n: Int): Char = ('a' + (ch - 'a' + n) % 26).toChar()