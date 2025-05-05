object Solution {
    def searchRange(nums: Array[Int], target: Int): Array[Int] = {
        def searchLeft(lo: Int, hi: Int, provisional: Int): Int = {
            if (lo > hi) return provisional
            val mi = lo + (hi - lo) / 2
            if (nums(mi) == target)
                searchLeft(lo, mi - 1, mi)
            else if (nums(mi) > target)
                searchLeft(lo, mi - 1, provisional)
            else
                searchLeft(mi + 1, hi, provisional)
        }
        def searchRight(lo: Int, hi: Int, provisional: Int): Int = {
            if (lo > hi) return provisional
            val mi = lo + (hi - lo) / 2
            if (nums(mi) == target)
                searchRight(mi + 1, hi, mi)
            else if (nums(mi) > target)
                searchRight(lo, mi - 1, provisional)
            else
                searchRight(mi + 1, hi, provisional)
        } 
        Array(searchLeft(0, nums.length - 1, -1), searchRight(0, nums.length - 1, -1))
    }
}