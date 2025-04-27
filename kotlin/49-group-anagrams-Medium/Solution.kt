class Solution {
    fun groupAnagrams(strs: Array<String>): List<List<String>> {
        return strs.groupBy{ it.sorted() }.toList().map{ it.second }
    }
}

private fun String.sorted() = String(toCharArray().apply{ sort() })