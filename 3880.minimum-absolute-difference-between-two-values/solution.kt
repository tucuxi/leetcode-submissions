class Solution {
    fun minAbsoluteDifference(nums: IntArray): Int {
        var i1: Int? = null
        var i2: Int? = null
        var minDiff = Int.MAX_VALUE

        fun minimize(i: Int?, j: Int) {
            if (i != null) {
                minDiff = minOf(j - i, minDiff)
            }
        }

        nums.forEachIndexed { i, num ->
            if (num == 1) {
                minimize(i2, i)
                i1 = i
            } else if (num == 2) {
                minimize(i1, i)
                i2 = i
            }        
        }

        return if (minDiff == Int.MAX_VALUE) -1 else minDiff
    }
}