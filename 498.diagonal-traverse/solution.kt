class Solution {
    fun findDiagonalOrder(mat: Array<IntArray>): IntArray {
        val m = mat.size
        val n = mat.first().size
        val sequence = sequence {
            var dr = -1
            var dc = 1
            var r = 0
            var c = 0
            repeat(m * n) {
                yield(mat[r][c])
                r += dr
                c += dc
                when {
                    r < 0 && c >= n -> {
                        dr = 1
                        dc = -1
                        r = 1
                        c = n - 1
                    }
                    r >= m && c < 0 -> {
                        dr = -1
                        dc = 1
                        r = m - 1
                        c = 1
                    }
                    r < 0 -> {
                        dr = 1
                        dc = -1
                        r = 0
                    }
                    r >= m -> {
                        dr = -1
                        dc = 1
                        r = m - 1
                        c += 2
                    }
                    c < 0 -> {
                        dr = -1
                        dc = 1
                        c = 0
                    }
                    c >= n -> {
                        dr = 1
                        dc = -1
                        r += 2
                        c = n - 1
                    }
                }
            }
        }
        return sequence.toList().toIntArray()    
    }
}