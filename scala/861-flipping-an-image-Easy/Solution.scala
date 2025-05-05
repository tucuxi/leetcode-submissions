object Solution {
    def flipAndInvertImage(A: Array[Array[Int]]): Array[Array[Int]] = {
        for (i <- 0 until A.length) {
            val jend = A(i).length - 1
            for (j <- 0 to jend / 2) {
                val h = A(i)(j)
                A(i)(j) = 1 - A(i)(jend - j)
                A(i)(jend - j) = 1 - h
            }
        }
        A
    }
}