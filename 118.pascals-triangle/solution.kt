class Solution {
    fun generate(numRows: Int): List<List<Int>> {
        val triangle = mutableListOf<List<Int>>()
        for (i in 0 until numRows) {
            val row = MutableList(i + 1) { j ->
                if (j == 0 || j == i) {
                    1
                } else {
                    triangle[i - 1][j] + triangle[i - 1][j - 1]
                }
            }
            triangle.add(row)
        }
        return triangle
    }
}