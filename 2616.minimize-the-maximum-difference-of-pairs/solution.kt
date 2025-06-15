class Solution {
    fun minimizeMax(nums: IntArray, p: Int): Int {
        val sortedNums = nums.sorted()

        fun canFormPairs(m: Int): Boolean {
            var count = 0
            var i = 1

            while (i <= sortedNums.lastIndex) {
                if (sortedNums[i] - sortedNums[i - 1] <= m) {
                    count++
                    i++
                }
                i++
            }
            return count >= p
        }

        var low = 0
        var high = sortedNums.last() - sortedNums.first()

        while (low < high) {
            val mid = low + (high - low) / 2
            if (canFormPairs(mid)) {
                high = mid
            } else {
                low = mid + 1
            }
        }
        return high
    }
}