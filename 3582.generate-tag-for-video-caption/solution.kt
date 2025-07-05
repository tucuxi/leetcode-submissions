class Solution {
    fun generateTag(caption: String): String {
        return '#' + caption
            .split(" ")
            .filter { s -> s.isNotEmpty() }
            .mapIndexed { i, s -> 
                val t = s.lowercase().replace("[^a-z]", "")
                if (i == 0) "$t" else t.first().uppercase() + t.drop(1)
            }
            .joinToString("")
            .take(99)
    }
}