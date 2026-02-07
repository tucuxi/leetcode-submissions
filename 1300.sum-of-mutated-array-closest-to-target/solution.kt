const val MAX_VALUE = 100000
class Solution {
    fun findBestValue(arr: IntArray, target: Int): Int {
        
        fun mutatedSum(value: Int) = arr.sumOf { minOf(it, value) }

        var l = 1
        var r = MAX_VALUE
        
        while (l < r) {
            val m = l + (r - l) / 2
            val sum = mutatedSum(m)
            if (sum < target) {
                l = m + 1
            } else {
                r = m
            }
        }

        return when {
            l == MAX_VALUE -> arr.max()
            target - mutatedSum(l - 1) <= mutatedSum(l) - target -> l - 1
            else -> l
        }
    }
}