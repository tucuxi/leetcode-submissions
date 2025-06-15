class Solution {
    fun wordSubsets(words1: Array<String>, words2: Array<String>): List<String> {
		val h2 = ('a'..'z').associate { ch -> ch to words2.maxOf { word -> word.count { it == ch } } }
        return words1.filter { w1 ->
            ('a'..'z').all { ch ->
                w1.count { it == ch } >= h2.getOrDefault(ch, 0)
            }
        }
    }
}