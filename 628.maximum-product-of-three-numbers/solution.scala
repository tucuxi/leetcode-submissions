object Solution {
    def maximumProduct(nums: Array[Int]): Int = {
        val sortedNums = nums.sorted
        val largest = sortedNums.takeRight(3)
        val smallest = sortedNums.take(2)
        val p = largest(0) * largest(1) * largest(2)
        val q = smallest(0) * smallest(1) * largest(2)
        math.max(p, q)
    }
}