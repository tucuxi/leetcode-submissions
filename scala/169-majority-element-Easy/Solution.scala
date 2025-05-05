object Solution {
    def majorityElement(nums: Array[Int]): Int = {
        val s = nums.sorted
        s(nums.length / 2)
    }
}