object Solution {
    def sortArrayByParity(A: Array[Int]): Array[Int] = {
        val (even, odd) = A.partition(_ % 2 == 0)
        even ++ odd
    }
}