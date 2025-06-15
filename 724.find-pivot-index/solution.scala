object Solution {
    def pivotIndex(nums: Array[Int]): Int = {
        def rec(index: Int, lsum: Int, rsum: Int): Int = {
            if (index == nums.length) -1
            else if (lsum == rsum - nums(index)) index
            else rec(index + 1, lsum + nums(index), rsum - nums(index))
        }
        rec(0, 0, nums.sum)
    }
}