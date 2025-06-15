object Solution {
    def removeDuplicates(nums: Array[Int]): Int = {
        if (nums.length == 0)
            0
        else {
            var last = 0
            for (i <- 1 until nums.length) {
                if (nums(i) != nums(last)) {
                    last += 1
                    nums(last) = nums(i)
                }
            }
            last + 1
        }
    }
}