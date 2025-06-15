import scala.collection.mutable.Queue

class RecentCounter() {

    private var q = Queue[Int]()

    def ping(t: Int): Int = {
        q = q.dropWhile(_ < t - 3000)
        q += t
        q.length
    }

}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = new RecentCounter()
 * var param_1 = obj.ping(t)
 */