import scala.collection.mutable.HashMap

object Solution {
    def findLHS(nums: Array[Int]): Int = {
        var h = HashMap[Int, Int]()
        for (num <- nums) {
            h(num) = h.getOrElse(num, 0) + 1
        }
        var m = 0
        for ((k, v) <- h) {
            val v1 = h.getOrElse(k + 1, 0)
            if (v1 > 0) m = Math.max(m, v + v1)
        }
        m
    }
}