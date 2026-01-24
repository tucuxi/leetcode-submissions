class Solution {
    fun numMatchingSubseq(s: String, words: Array<String>): Int {
        val occurrences = Array(26) { TreeSet<Int>() }
        s.forEachIndexed { i, c -> occurrences[c - 'a'].add(i) }

        return words.count { w ->
            var i = -1
            w.all { c ->
                occurrences[c - 'a']
                    .higher(i)
                    ?.let { i = it; true }
                    ?: false
            }
        }
    }
}