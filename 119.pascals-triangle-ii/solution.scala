object Solution {
    def getRow(rowIndex: Int): List[Int] = {
        if (rowIndex == 0)
            List(1)
        else {
            val prev = getRow(rowIndex - 1)
            val inner = for (i <- 0 until prev.length - 1) yield prev(i) + prev(i + 1)
            (1 :: inner.toList) :+ 1
        }
    }
}