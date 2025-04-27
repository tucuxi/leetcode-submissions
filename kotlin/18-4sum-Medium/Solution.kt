class Solution {
    fun fourSum(nums: IntArray, target: Int): List<List<Int>> {
        nums.sort()
        val res: MutableSet<List<Int>> = mutableSetOf()
        for (a in 0 .. nums.lastIndex - 3) {
            for (b in a + 1 .. nums.lastIndex - 2) {
                var c = b + 1
                var d = nums.lastIndex
                while (c < d) {
                    val sum = nums[a] + nums[b] + nums[c] + nums[d]
                    if (sum < target) {
                        c++
                    } else if (sum > target) {
                        d--
                    } else {
                        res.add(listOf(nums[a], nums[b], nums[c], nums[d]))
                        c++
                        d--
                    }
                }
            }
        }
        return res.toList()
    }
}