class Solution {
    fun commonChars(A: Array<String>): List<String> {
        if (A.isEmpty()) return emptyList()
        val res = A.drop(1).fold(normalize(A[0])) {
            acc, elem -> intersect(acc, normalize(elem))
        }
        return format(res)    
    }

    fun normalize(s: String): Map<Char, Int> =
        s.groupingBy { it }.eachCount()
    
    fun intersect(a: Map<Char, Int>, b: Map<Char, Int>): Map<Char, Int> =
        a.keys.associate {
            it to Math.min(a.getOrDefault(it, 0), b.getOrDefault(it, 0))
        }
    
    fun format(a: Map<Char, Int>) =
        a.flatMap { entry -> List(entry.value) { entry.key.toString() } }
}