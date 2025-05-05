object Solution {
    def sortedSquares(A: Array[Int]): Array[Int] = {
        A.map(x => x * x).sorted
    }
}