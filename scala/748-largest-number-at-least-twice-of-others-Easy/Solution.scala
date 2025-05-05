object Solution {
    def dominantIndex(nums: Array[Int]): Int = {
        var iMax = 0
        for (i <- 1 until nums.length)
            if (nums(i) > nums(iMax)) iMax = i
        for (i <- 0 until nums.length if i != iMax)
            if (2 * nums(i) > nums(iMax))
                return -1
        iMax
    }
}