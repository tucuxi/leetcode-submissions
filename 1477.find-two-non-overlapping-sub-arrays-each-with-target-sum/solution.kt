class Solution {
    fun minSumOfLengths(arr: IntArray, target: Int): Int {
        val n = arr.size
        val dp = IntArray(n) { Int.MAX_VALUE }
        
        var minSumOfLengths = Int.MAX_VALUE
        var currentSum = 0
        var left = 0

        for (right in 0 until n) {
            currentSum += arr[right]

            while (currentSum > target) {
                currentSum -= arr[left]
                left++
            }
            
            if (currentSum == target) {
                val currentLength = right - left + 1
                
                if (left > 0 && dp[left - 1] != Int.MAX_VALUE) {
                    minSumOfLengths = min(minSumOfLengths, currentLength + dp[left - 1])
                }
                
                dp[right] = currentLength
            }
            
            if (right > 0) {
                dp[right] = min(dp[right], dp[right - 1])
            }
        }
        
        return if (minSumOfLengths == Int.MAX_VALUE) -1 else minSumOfLengths
    }
}