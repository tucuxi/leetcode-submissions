import scala.collection.mutable.{HashMap, ArrayBuffer}

object Solution {
    def relativeSortArray(arr1: Array[Int], arr2: Array[Int]): Array[Int] = {
        val h = HashMap[Int, Int]()
        for (k <- arr1) h(k) = h.getOrElse(k, 0) + 1
        val res = ArrayBuffer[Int]()
        for (k <- arr2) {
            val v = h.remove(k).get
            for (n <- 0 until v) res += k
        }
        for (k <- h.keys.toList.sorted; n <- 0 until h(k)) res += k
        res.toArray
    }
}