object Solution {
    def search(nums: Array[Int], target: Int): Int = {

        def bsearch(lo: Int, hi: Int): Int = {
            if (lo > hi) return -1
            val pv = lo + (hi - lo) / 2
            if (nums(pv) == target)
                pv
            else if (nums(pv) > target)
                bsearch(lo, pv - 1)
            else
                bsearch(pv + 1, hi)
        }
        
        def shouldIntervalContainTarget(lo: Int, hi: Int): Boolean = {
            nums(lo) <= target && target <= nums(hi)    
        }
        
        def isIntervalOrdered(lo: Int, hi: Int): Boolean = {
            nums(lo) <= nums(hi)    
        }
        
        def rsearch(lo: Int, hi: Int): Int = {
            if (lo > hi) return -1
            val pv = lo + (hi - lo) / 2
            if (isIntervalOrdered(lo, pv)) {
                if (shouldIntervalContainTarget(lo, pv)) bsearch(lo, pv)
                else rsearch(pv + 1, hi)
            } else {
                if (shouldIntervalContainTarget(pv, hi)) bsearch(pv, hi)
                else rsearch(lo, pv - 1)
            }
        }

        rsearch(0, nums.length - 1)        
    }
}