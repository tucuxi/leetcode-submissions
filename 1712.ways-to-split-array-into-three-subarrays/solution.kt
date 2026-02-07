const val MOD = 1_000_000_007

class Solution {
    fun waysToSplit(nums: IntArray): Int {
        val n = nums.size
        val prefix = IntArray(n)

        prefix[0] = nums[0]
        for (k in 1 until n) {
            prefix[k] = prefix[k - 1] + nums[k]
        }
        
        val totalSum = prefix[n - 1]

        fun lowerBound(l: Int, r: Int, target: Int): Int {
            var low = l
            var high = r
            var ans = r + 1
            
            while (low <= high) {
                val mid = low + (high - low) / 2
                if (prefix[mid] >= target) {
                    ans = mid
                    high = mid - 1
                } else {
                    low = mid + 1
                }
            }
            return ans
        }

        fun upperBound(l: Int, r: Int, target: Int): Int {
            var low = l
            var high = r
            var ans = l - 1
            
            while (low <= high) {
                val mid = low + (high - low) / 2
                if (prefix[mid] <= target) {
                    ans = mid
                    low = mid + 1
                } else {
                    high = mid - 1
                }
            }
            return ans
        }

        var res = 0
        
        for (i in 0 until n - 2) {
            val leftSum = prefix[i]
            val lowerJ = lowerBound(i + 1, n - 2, 2 * leftSum)
            val upperJ = upperBound(i + 1, n - 2, (totalSum + leftSum) / 2)
            
            if (upperJ >= lowerJ) {
                res = (res + upperJ - lowerJ + 1) % MOD
            }
        }
        return res
    }
}