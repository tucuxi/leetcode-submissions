class Solution {
    fun countLargestGroup(n: Int): Int {
        return (1..n)
            .groupingBy { digitsSum(it) }
            .eachCount()
            .values
            .groupingBy { it } 
            .eachCount()
            .maxBy { it.key }
            .value
    }

    tailrec fun digitsSum(n: Int, acc: Int = 0): Int {
        return if (n == 0) {
            acc
        } else {
            digitsSum(n / 10, acc + n % 10)
        }
    }
}