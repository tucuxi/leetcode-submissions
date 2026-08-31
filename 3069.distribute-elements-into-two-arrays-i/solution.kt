class Solution {
    fun resultArray(nums: IntArray): IntArray {
        val arr1 = mutableListOf(nums[0])
        val arr2 = mutableListOf(nums[1])

        for (i in 2..nums.lastIndex) {
            if (arr1.last() > arr2.last()) {
                arr1 += nums[i]
            } else {
                arr2 += nums[i]
            }
        }

        val res = arr1 + arr2
        return res.toIntArray()
    }
}