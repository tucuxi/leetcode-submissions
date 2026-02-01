class Solution {
    fun longestSubsequence(nums: IntArray): Int {
        var maxLen = 0
        val tails = IntArray(nums.size)

        for (bit in 0 until 30) {
            var size = 0
            for (num in nums) {
                if ((num shr bit) and 1 == 1) {
                    var left = 0
                    var right = size
                    while (left < right) {
                        val mid = (left + right) / 2
                        if (tails[mid] < num) {
                            left = mid + 1
                        } else {
                            right = mid
                        }
                    }
                    if (left == size) {
                        tails[size++] = num
                    } else {
                        tails[left] = num
                    }
                }
            }
            maxLen = maxOf(size, maxLen)
        }

        return maxLen
    }
}
