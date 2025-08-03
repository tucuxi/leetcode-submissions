class Solution {
    fun generate(numRows: Int): List<List<Int>> =
        generateSequence(listOf(1)) { row ->
            List(row.size + 1) { i ->
                row.getOrElse(i - 1) { 0 } + row.getOrElse(i) { 0 }
             }
        }
        .take(numRows)
        .toList()
}