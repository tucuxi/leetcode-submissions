class Solution {
    fun largeGroupPositions(s: String): List<List<Int>> {
        val res = mutableListOf<List<Int>>()
        var a = 0
        for (i in s.indices) {
            if (s[i] != s[a]) {
                res.addIfLarge(a, i - 1)
                a = i
            }   
        }
        res.addIfLarge(a, s.lastIndex)
        return res
    }
}

inline fun MutableList<List<Int>>.addIfLarge(start: Int, end: Int) {
    if (end - start >= 2) {
        add(listOf(start, end))
    }
}