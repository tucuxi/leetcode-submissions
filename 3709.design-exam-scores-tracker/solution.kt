class ExamTracker() {

    private val prefix = TreeMap<Int, Long>()
    private var total = 0L

    fun record(time: Int, score: Int) {
        total += score
        prefix[time] = total

    }

    fun totalScore(startTime: Int, endTime: Int): Long {
        val h = prefix.floorEntry(endTime)?.value ?: 0
        val l = prefix.lowerEntry(startTime)?.value ?: 0
        return h - l
    }

}

/**
 * Your ExamTracker object will be instantiated and called as such:
 * var obj = ExamTracker()
 * obj.record(time,score)
 * var param_2 = obj.totalScore(startTime,endTime)
 */