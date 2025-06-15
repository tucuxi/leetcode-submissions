object Solution {
    def fib(N: Int): Int = {
        val f = new Array[Int](Math.max(N + 1, 2))
        f(0) = 0
        f(1) = 1
        for (i <- 2 to N) f(i) = f(i - 1) + f(i - 2)
        f(N)
    }
}