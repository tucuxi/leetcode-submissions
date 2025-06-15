const val marker = '*'

class Solution {
    fun clearStars(s: String): String {
        val indices = Array(26) { ArrayDeque<Int>() }
        val res = s.toCharArray()
        res.forEachIndexed { i, c ->
            if (c == marker) {
                val j = indices
                    .first { it.isNotEmpty() }
                    .removeLast()
                res[j] = marker
            } else {
                indices[c - 'a'].add(i)
            }
        }
        return buildString {
            res.forEach { c ->
                if (c != marker) {
                    append(c)
                }
            }
        }
    }
}