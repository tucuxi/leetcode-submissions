object Solution {
    def ways(a: Int, b: Int): Stream[Int] = a #:: ways(b, a + b)

    def climbStairs(n: Int): Int = {
        ways(1, 2)(n - 1)
    }
}