object Solution {
    def isMonotonic(A: Array[Int]): Boolean = {
        var inc = 0
        var dec = 0
        for (i <- 1 until A.length) {
            val d = A(i) - A(i - 1)
            if (d >= 0) inc += 1
            if (d <= 0) dec += 1
        }
        Math.max(inc, dec) >= A.length - 1
    }
}