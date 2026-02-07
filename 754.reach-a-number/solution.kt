class Solution {
    fun reachNumber(target: Int): Int {
        val absTarget = abs(target)
        var sum = 0
        var k = 0
        while (sum < absTarget) {
            sum += ++k
        }
        return if ((sum - absTarget) % 2 == 0) k else k + 1 + k % 2 
    }
}