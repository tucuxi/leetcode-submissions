object Solution {
    def findMin(nums: Array[Int]): Int = {
        def find(lo: Int, hi: Int): Int = {
            if (nums(lo) <= nums(hi))
                nums(lo)
            else {
                val pv = lo + (hi - lo) / 2
                if (nums(pv) < nums(hi))
                    find(lo, pv)
                else
                    find(pv + 1, hi)
            }
        }
        find(0, nums.length - 1)
    }
}