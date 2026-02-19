class Solution {
    fun prefixConnected(words: Array<String>, k: Int): Int {
        val prefixCounts = mutableMapOf<String, Int>()

		words
			.filter { it.length >= k }
			.forEach { word ->
                val prefix = word.substring(0, k)
                prefixCounts[prefix] = prefixCounts.getOrDefault(prefix, 0) + 1
			}

        return  prefixCounts.values.count { it >= 2 }
    }
}