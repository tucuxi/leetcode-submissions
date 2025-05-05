object Solution {
    def findLengthOfLCIS(nums: Array[Int]): Int = {
        if (nums.isEmpty) return 0
        var k = 1
        var m = 1
        for (i <- 1 until nums.length)
            if (nums(i) > nums(i - 1)) {
                k += 1
                if (k > m) m = k
            } else {
                k = 1
            }
        m
    }
}