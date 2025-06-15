import scala.collection.mutable.ListBuffer

object Solution {
    def threeSum(nums: Array[Int]): List[List[Int]] = {
        val s = nums.sorted
        val r = s.length - 1
        val res = ListBuffer[List[Int]]()
        for (i <- 0 to r - 2 if i == 0 || s(i) != s(i -1)) {
            var j = i + 1
            var k = r
            while (j < k) {
                val sum = s(i) + s(j) + s(k)
                if (sum < 0) {
                    j += 1
                } else if (sum > 0) {
                    k -= 1
                } else {
                    res += List(s(i), s(j), s(k))
                    j += 1
                    while (j < r && s(j) == s(j - 1)) j += 1
                    while (k > j && s(k) == s(k - 1)) k -= 1
                }
            }
        }
        res.toList
    }
}