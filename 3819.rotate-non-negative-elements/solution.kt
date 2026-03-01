class Solution {
    fun rotateElements(nums: IntArray, k: Int): IntArray {
        val p = nums.filter { it >= 0 }
        
        if (p.isEmpty()) {
            return nums
        }
        
        var j = k
        
        nums.indices.forEach { i ->
            if (nums[i] >= 0) {
                nums[i] = p[j++ % p.size]
            }
        }
        return nums
    }
}