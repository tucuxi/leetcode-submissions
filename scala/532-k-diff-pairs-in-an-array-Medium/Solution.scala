object Solution {
    def findPairs(nums: Array[Int], k: Int): Int = {
        if (k < 0) {
            0
        } else if (k == 0) {
            nums.groupBy(identity).count(_._2.length > 1)
        }
        else {
            val numSet = nums.toSet
            numSet.count(n => numSet.contains(n + k))
        }
    }
}