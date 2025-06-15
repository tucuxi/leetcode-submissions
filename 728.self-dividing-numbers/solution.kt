class Solution {
    fun selfDividingNumbers(left: Int, right: Int) =
        (left..right).filter { selfDividingNumber(it) }
    
    fun selfDividingNumber(num: Int) =
        "$num".all { it != '0' && num % (it - '0').toInt() == 0 }
}