class Solution {
    fun twoOutOfThree(nums1: IntArray, nums2: IntArray, nums3: IntArray): List<Int> {
        val map = mutableMapOf<Int, Int>()
        add(map, nums1)
        add(map, nums2)
        add(map, nums3)
        return map.filter { (num, count) -> count > 1 }.map { (num, count) -> num }
    }
    
    fun add(map: MutableMap<Int, Int>, nums: IntArray) {
        for (num in nums.distinct()) {
            map[num] = map.getOrDefault(num, 0) + 1
        }
    }
}