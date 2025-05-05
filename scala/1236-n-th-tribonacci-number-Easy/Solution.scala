object Solution {
    def tribonacci(n: Int): Int = {
        var t = new Array[Int](Math.max(3, n + 1))
        t(0) = 0
        t(1) = 1
        t(2) = 1
        for (i <- 3 to n)
            t(i) = t(i - 3) + t(i - 2) + t(i - 1)
        t(n)
    }
}