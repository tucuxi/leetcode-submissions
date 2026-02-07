class Solution {
    fun minRemoval(nums: IntArray, k: Int): Int {
        var i = 0
        var j = 0
        var l = 0

        nums.sort()

        while (j < nums.size) {
            if (nums[j] <= k * nums[i].toLong()) {
                j++
                l = maxOf(l, j - i)
            } else {
                i++
            }
        }

        return nums.size - l
    }
}