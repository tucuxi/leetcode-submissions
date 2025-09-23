class Solution {
    fun evaluate(s: String, knowledge: List<List<String>>): String {
        val map = knowledge.associate { (k, v) -> k to v }
        return buildString {
            var j = -1
            s.forEachIndexed { i, ch ->
                when (ch) {
                    '(' -> {
                        j = i
                    }
                    ')' -> {
                        val k = s.substring(j + 1, i)
                        append(map.getOrDefault(k, "?"))
                        j = -1
                    }
                    else -> {
                        if (j == -1) {
                            append(ch)
                        }
                    }
                }
            }
        }
    }
}