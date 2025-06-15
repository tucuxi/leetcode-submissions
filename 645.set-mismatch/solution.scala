import scala.collection.mutable.BitSet

object Solution {
    def findErrorNums(nums: Array[Int]): Array[Int] = {
        val bs = new BitSet(nums.length + 1)
        var dp = 0
        for (n <- nums)
            if (!bs.add(n)) dp = n
        for (i <- 1 to nums.length)
            if (!bs.contains(i))
                return Array(dp, i)
        Array(dp, 0)
    }
}