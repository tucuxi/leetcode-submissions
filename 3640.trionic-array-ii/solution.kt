class Solution {
    fun maxSumTrionic(nums: IntArray): Long {
        val n = nums.size
        if (n < 4) return Long.MIN_VALUE 

        val lInc = LongArray(n)
        val rInc = LongArray(n)
        val pref = LongArray(n + 1)

        lInc[0] = nums[0].toLong()
        for (i in 1 until n) {
            lInc[i] = nums[i].toLong() + if (nums[i] > nums[i - 1]) maxOf(0L, lInc[i - 1]) else 0L
        }

        rInc[n - 1] = nums[n - 1].toLong()
        for (i in n - 2 downTo 0) {
            rInc[i] = nums[i].toLong() + if (nums[i] < nums[i + 1]) maxOf(0L, rInc[i + 1]) else 0L
        }

        for (i in 0 until n) pref[i + 1] = pref[i] + nums[i]

        var maxTrionicSum = Long.MIN_VALUE
        var i = 0

        while (i < n - 1) {
            if (nums[i] > nums[i + 1]) {
                val start = i
                while (i < n - 1 && nums[i] > nums[i + 1]) {
                    i++
                }
                val end = i

                val isValidPeak = start > 0 && nums[start] > nums[start - 1]
                val isValidValley = end < n - 1 && nums[end] < nums[end + 1]

                if (isValidPeak && isValidValley) {
                    val leftPart = lInc[start - 1]
                    val middlePart = pref[end + 1] - pref[start]
                    val rightPart = rInc[end + 1]
                    
                    maxTrionicSum = maxOf(maxTrionicSum, leftPart + middlePart + rightPart)
                }
            } else {
                i++
            }
        }
        
        return maxTrionicSum
    }
}