class Solution {
    fun minimumK(nums: IntArray): Int {
        var left = 1
        var right = 100000
        var ans = right

        while (left <= right) {
            val mid = left + (right - left) / 2            
            var operations = 0L
            for (num in nums) {
                operations += (num + mid - 1) / mid
            }
            if (operations <= mid.toLong() * mid) {
                ans = mid
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
        
        return ans
    }
}
