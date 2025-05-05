object Solution {
    def maxSubArray(nums: Array[Int]): Int = {
        var largest = Int.MinValue
        var running = 0
        for (num <- nums) {
            running += num
            largest = largest.max(running)
            running = running.max(0)
        }
        largest
    }
}