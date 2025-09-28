class Solution {
    fun triangleNumber(nums: IntArray): Int {
        var res = 0
        nums.sort()
        for (i in nums.indices) {
            if (nums[i] == 0) continue
            var k = i + 2
            for (j in i + 1 until nums.lastIndex) {
                k = binarySearch(nums, k, nums.lastIndex, nums[i] + nums[j])
                res += k - j - 1
            }
        }
        return res
    }

    fun binarySearch(nums: IntArray, left: Int, right: Int, x: Int): Int {
        var l = left
        var r = right
        while (l <= r && r < nums.size) {
            val mid = (l + r) / 2
            if (nums[mid] >= x) {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
        return l
    }
}