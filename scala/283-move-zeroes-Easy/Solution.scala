object Solution {
    def moveZeroes(nums: Array[Int]): Unit = {
        var k = 0
        for (i <- 0 until nums.length) {
            if (nums(i) != 0) {
                nums(k) = nums(i)
                k += 1
            }
        }
        while (k < nums.length) {
            nums(k) = 0
            k += 1
        }
    }
}