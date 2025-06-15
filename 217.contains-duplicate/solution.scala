object Solution {
    def containsDuplicate(nums: Array[Int]): Boolean = {
        val orderedNums = nums.sorted
        for (i <- 1 until orderedNums.length)
            if (orderedNums(i - 1) == orderedNums(i))
                return true
        return false
    }
}