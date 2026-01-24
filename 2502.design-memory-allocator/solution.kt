class Allocator(n: Int) {
    // Stores free blocks: Start Index -> Size
    private val available = TreeMap<Int, Int>()
    // Stores allocated blocks: mID -> List of Blocks
    private val allocated = mutableMapOf<Int, MutableList<Pair<Int, Int>>>()

    init {
        available[0] = n
    }

    fun allocate(size: Int, mID: Int): Int {
        for ((start, blockSize) in available) {
            if (blockSize >= size) {
                // Found the leftmost block
                available.remove(start)
                if (blockSize > size) {
                    available[start + size] = blockSize - size
                }
                
                allocated.getOrPut(mID) { mutableListOf() }.add(start to size)
                return start
            }
        }
        return -1
    }

    fun freeMemory(mID: Int): Int {
        val blocksToFree = allocated.remove(mID) ?: return 0
        var totalFreedSize = 0

        for ((start, size) in blocksToFree) {
            totalFreedSize += size
            insertAndMerge(start, size)
        }
        return totalFreedSize
    }

    private fun insertAndMerge(start: Int, size: Int) {
        var newStart = start
        var newSize = size

        // 1. Check if there's a neighbor to the RIGHT
        val nextStart = available.ceilingKey(newStart + newSize)
        if (nextStart != null && nextStart == newStart + newSize) {
            newSize += available.remove(nextStart)!!
        }

        // 2. Check if there's a neighbor to the LEFT
        val prevStart = available.floorKey(newStart)
        if (prevStart != null && prevStart + available[prevStart]!! == newStart) {
            val prevSize = available.remove(prevStart)!!
            newStart = prevStart
            newSize += prevSize
        }

        // 3. Insert the newly merged block
        available[newStart] = newSize
    }
}