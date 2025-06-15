class Solution {
    fun summaryRanges(nums: IntArray): List<String> {
        val res = mutableListOf<String>()
        if (nums.size == 0) {
            return res
        }
        var i = 0
        var k = 0
        while (++k + i < nums.size) {
            if (nums[i + k] != nums[i] + k) {
                addRange(nums[i], k - 1, res)
                i += k
                k = 0
            }
        }
        if (i < nums.size) {
            addRange(nums[i], k - 1, res)
        }
        return res
    }
    
    fun addRange(a: Int, offset: Int, res: MutableList<String>) {
        if (offset == 0) {
            res.add("$a")
        } else {
            res.add("$a->${a + offset}")
        }
    }
}