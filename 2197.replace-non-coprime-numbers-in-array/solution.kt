class Solution {
    fun replaceNonCoprimes(nums: IntArray): List<Int> {
        val stack = ArrayDeque<Int>()
        nums.forEach { num ->
            var current = num
            while (stack.isNotEmpty()) {
                val gcd = gcd(current, stack.last())
                if (gcd == 1) break
                current = current / gcd * stack.removeLast()
            }
            stack.add(current)
        }
        return stack.toList()
    }

    tailrec fun gcd(x: Int, y: Int): Int = if (y == 0) x else gcd(y, x % y)
}