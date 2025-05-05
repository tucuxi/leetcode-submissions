object Solution {
    def singleNumber(nums: Array[Int]): Int = {
        nums.fold(0)(_ ^ _)
    }
}