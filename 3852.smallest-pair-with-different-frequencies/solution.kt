class Solution {
    fun minDistinctFreqPair(nums: IntArray): IntArray {
        val h = IntArray(101)
        nums.forEach { num -> h[num]++ }
        for (i in 1..h.lastIndex) {
            if (h[i] == 0) continue
            for (j in i+1..h.lastIndex) {
                if (h[j] != 0 && h[j] != h[i]) {
                    return intArrayOf(i, j)
                }
            }
        }
        return intArrayOf(-1, -1)
    }
}