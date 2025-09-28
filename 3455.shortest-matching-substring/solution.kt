class Solution {
    fun shortestMatchingSubstring(s: String, p: String): Int {
        val (part1, part2, part3) = p.split("*")
        val indices1 = findIndices(s, part1)
        val indices2 = findIndices(s, part2)
        val indices3 = if (part3.isEmpty()) {
            IntArray(s.length + 1) { it - 1 }.asList()
        } else {
            sequence {
                var index = s.indexOf(part3)
                while (index != -1) {
                    yield(index + part3.length - 1)
                    index = s.indexOf(part3, index + 1)
                }
            }.toList()
        }
        
        var minLength = Int.MAX_VALUE
        var p1 = 0
        var p3 = 0

        for (j in indices2) {
            while (p1 < indices1.size && indices1[p1] <= j - part1.length) {
                p1++
            }
            if (p1 == 0) continue
            val i = indices1[p1 - 1]

            while (p3 < indices3.size && indices3[p3] < j + part2.length - 1 + part3.length) {
                p3++
            }
            if (p3 == indices3.size) break
            val k = indices3[p3]

            minLength = minOf(minLength, k - i + 1)
        }
        return if (minLength == Int.MAX_VALUE) {
            -1
        } else {
            minLength
        }
    }

     private fun findIndices(s: String, part: String): List<Int> {
        return if (part.isEmpty()) {
            IntArray(s.length + 1) { it }.asList()
        } else {
            sequence {
                var index = s.indexOf(part)
                while (index != -1) {
                    yield(index)
                    index = s.indexOf(part, index + 1)
                }
            }.toList()
        }
    }
}