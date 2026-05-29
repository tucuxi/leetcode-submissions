class Solution {
    fun maxPoints(technique1: IntArray, technique2: IntArray, k: Int): Long {
        val n = technique1.size
        val delta = IntArray(n) { technique2[it] - technique1[it] }

        return (0 until n)
            .sortedBy { delta[it] }
            .withIndex()
            .sumOf { (i, j) ->
                if (i < k || delta[j] < 0) {
                    technique1[j].toLong()
                } else {
                    technique2[j].toLong()
                }
            }
    }
}