class Solution {
    fun maxProductPair(nums: IntArray, target: Int): IntArray {
        var i1 = -1
        var j1 = -1
        var p1 = Int.MIN_VALUE

        for (i in nums.indices) {
            val wanted = target - nums[i]
            if (wanted >= nums[i]) continue
            for (j in nums.indices) {
                if (j == i || nums[j] != wanted) continue
                val p = nums[i] * nums[j]
                if (p > p1) {
                    i1 = i
                    j1 = j
                    p1 = p
                }  
            }
        }

        return intArrayOf(i1, j1)
    }
}