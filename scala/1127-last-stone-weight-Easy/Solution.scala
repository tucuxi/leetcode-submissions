import scala.collection.mutable.PriorityQueue

object Solution {
    def lastStoneWeight(stones: Array[Int]): Int = {
        val q = PriorityQueue[Int]()
        q ++= stones
        while (q.length > 1) {
            val y = q.dequeue
            val x = q.dequeue
            if (x != y) q.enqueue(y - x)
        }
        if (q.isEmpty) 0 else q.dequeue
    }
}