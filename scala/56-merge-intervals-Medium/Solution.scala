object Solution {
    def merge(intervals: Array[Array[Int]]): Array[Array[Int]] = {
        intervals
            .sortBy(_(0))
            .foldLeft(List.empty[Array[Int]])((seq, interval) =>
                if (seq.isEmpty || interval(0) > seq.last(1)) {
                    seq :+ interval
                } else {
                    seq.last(1) = Math.max(seq.last(1), interval(1))
                    seq
                }
            )
            .toArray
    }
}