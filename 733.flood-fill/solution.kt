class Solution {
    fun floodFill(image: Array<IntArray>, sr: Int, sc: Int, newColor: Int): Array<IntArray> {
        val directions = listOf(0 to -1, -1 to 0, 0 to 1, 1 to 0)
        val m = image.size
        val n = image[0].size
        val color = image[sr][sc]
        
        fun recurse(r: Int, c: Int) {
            image[r][c] = newColor
            for (dir in directions) {
                val nextR = r + dir.first
                val nextC = c + dir.second
                if (nextR in 0 until m && nextC in 0 until n && image[nextR][nextC] == color) {
                    recurse(nextR, nextC)
                }
            }
        }
        
        if (newColor != color) {        
            recurse(sr, sc)
        }
        return image
    }
}