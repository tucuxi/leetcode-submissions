object Solution {
    def numEquivDominoPairs(dominoes: Array[Array[Int]]): Int = {
        val count = new Array[Int](100)
        for (d <- dominoes) {
            val n = if (d(0) <= d(1)) d(0) + 10 * d(1) else 10 * d(0) + d(1)
            count(n) += 1
        }
        count.fold(0) {(z, i) => z + i * (i - 1) / 2}
    }
}