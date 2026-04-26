class Solution {
    fun maxNonOverlapping(nums: IntArray, target: Int): Int {
        var res = 0
        var sum = 0
        val vis = mutableSetOf(0)

        for (num in nums) {
            sum += num
            if (sum - target in vis) {
                res++
                sum = 0
                vis.clear()
            }
            vis += sum
        }

        return res
    }
}