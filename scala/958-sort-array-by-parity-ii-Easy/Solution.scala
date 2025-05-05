object Solution {
    def sortArrayByParityII(A: Array[Int]): Array[Int] = {
        val p = A.partition(_ % 2 == 0)
        val s = for {
            (even, odd) <- p._1 zip p._2
            r <- List(even, odd)
        } yield r
        s.toArray
    }
}