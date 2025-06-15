object Solution {
    def uniqueOccurrences(arr: Array[Int]): Boolean = {
        val ns = arr.groupBy(identity).map(_._2.length)
        ns.toList.length == ns.toSet.size
    }
}