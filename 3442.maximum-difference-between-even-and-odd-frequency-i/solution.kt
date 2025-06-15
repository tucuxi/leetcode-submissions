class Solution {
    fun maxDifference(s: String): Int {
        val (evenFreqList, oddFreqList) = s
            .groupingBy { it }
            .eachCount()
            .values
            .partition { it % 2 == 0 }
        return oddFreqList.max() - evenFreqList.min()
    }
}