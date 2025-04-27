class Solution {
    fun fizzBuzz(n: Int): List<String> =
        (1..n).map { i ->
            when {
                d(i, 15) -> "FizzBuzz"
                d(i, 3) -> "Fizz"
                d(i, 5) -> "Buzz"
                else -> "${i}"
            }
        }
}
    
inline fun d(a: Int, b: Int) = a % b == 0