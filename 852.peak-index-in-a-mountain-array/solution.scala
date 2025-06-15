object Solution {
    def peakIndexInMountainArray(A: Array[Int]): Int = {
        var lo = 0
        var hi = A.length - 1
        while (lo <= hi) {
            val mi = lo + (hi -lo) / 2
            if (A(mi - 1) < A(mi)) {
                if (A(mi) > A(mi + 1)) return mi
                lo = mi
            } else
                hi = mi
        }
        -1
    }
}