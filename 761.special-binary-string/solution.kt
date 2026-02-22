class Solution {
    fun makeLargestSpecial(s: String): String {
        if (s.isEmpty()) return s

        val components = mutableListOf<String>()
        var count = 0
        var start = 0

        for (i in s.indices) {
            count += if (s[i] == '1') 1 else -1
            if (count == 0) {
                val inner = makeLargestSpecial(s.substring(start + 1, i))
                components.add("1${inner}0")
                start = i + 1
            }
        }

        return components.sortedDescending().joinToString("")
    }
}