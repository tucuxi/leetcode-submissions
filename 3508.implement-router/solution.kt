class Router(val memoryLimit: Int) {

    data class DataPacket(val source:Int, val destination: Int, val timestamp: Int) {
        fun toIntArray() = intArrayOf(source, destination, timestamp)
    }
  
    private val packetsQueue = ArrayDeque<DataPacket>(memoryLimit)

    private val packetsSet = mutableSetOf<DataPacket>()

    private val timestampsByDestination = mutableMapOf<Int, MutableList<Int>>()

    fun addPacket(source: Int, destination: Int, timestamp: Int): Boolean {
        val packet = DataPacket(source, destination, timestamp)

        if (!packetsSet.add(packet)) return false

        if (packetsQueue.size == memoryLimit) {
            forwardPacket()
        }

        packetsQueue.addLast(packet)
        timestampsByDestination.getOrPut(destination) { mutableListOf() }.add(timestamp)
        return true
    }

    fun forwardPacket(): IntArray {
        if (packetsQueue.isEmpty()) return intArrayOf()
        
        val packet = packetsQueue.removeFirst()
        packetsSet.remove(packet)
        timestampsByDestination.getValue(packet.destination).removeAt(0)
        return packet.toIntArray()
    }

    fun getCount(destination: Int, startTime: Int, endTime: Int): Int {
        val timestamps = timestampsByDestination[destination] ?: return 0
        return bisect(timestamps, endTime + 1) - bisect(timestamps, startTime)
    }

    companion object {
        private fun bisect(timestamps: List<Int>, target: Int): Int {
            var l = 0
            var r = timestamps.size
            while (l < r) {
                val m = (l + r) / 2
                if (timestamps[m] < target) {
                    l = m + 1
                } else {
                    r = m
                }
            }
            return l
        }
    }
}

/**
 * Your Router object will be instantiated and called as such:
 * var obj = Router(memoryLimit)
 * var param_1 = obj.addPacket(source,destination,timestamp)
 * var param_2 = obj.forwardPacket()
 * var param_3 = obj.getCount(destination,startTime,endTime)
 */