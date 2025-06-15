object Solution {
    def arrayPairSum(nums: Array[Int]): Int = {
        nums.sorted.grouped(2).map(_.min).sum
    }
}