object Solution {
    def reverse(x: Int): Int = {
        def revAcc(acc: Int, x: Int): Int = {
            if (x == 0)
                acc
            else if (acc > Int.MaxValue / 10 || acc < Int.MinValue / 10)
                0
            else
                revAcc(acc * 10 + x % 10, x / 10)
        }
        revAcc(0, x)
    }
}