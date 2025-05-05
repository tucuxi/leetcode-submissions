object Solution {
    def twoSum(numbers: Array[Int], target: Int): Array[Int] = {
        def twoSumRec(head: Int, tail: Int): (Int, Int) = {
            if (head >= tail) return (0, 0)
            val sum = numbers(head) + numbers(tail)
            if (sum < target)
                twoSumRec(head + 1, tail)
            else if (sum > target)
                twoSumRec(head, tail - 1)
            else
                (head, tail)
        }
        val (head, tail) = twoSumRec(0, numbers.length - 1)
        Array(head + 1, tail + 1)
    }
}