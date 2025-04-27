class Solution {
    fun subsetsWithDup(nums: IntArray): List<List<Int>> =
        with (mutableSetOf<List<Int>>(emptyList())) {
            nums.sorted().forEach { n ->
                addAll(map { elem -> elem + n })
            }
            toList()
        }
}