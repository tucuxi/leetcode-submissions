private const val MOD = 1_000_000_007

class Solution {
    fun countPermutations(complexity: IntArray): Int =
        complexity.drop(1).let { rest ->
            if (rest.any { it <= complexity[0] }) 0 else modFac(rest.size)
        } 
        
    tailrec fun modFac(n: Int, acc: Long = 1): Int =
        if (n == 1) acc.toInt() else modFac(n - 1, n * acc % MOD)
}