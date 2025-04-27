class Solution {
    fun maximumUnits(boxTypes: Array<IntArray>, truckSize: Int): Int {
        var loadedUnits = 0
        var remainingSpace = truckSize
        for (box in boxTypes.sortedByDescending { it[1]} ) {
            val numberOfBoxes = Math.min(box[0], remainingSpace)
            loadedUnits += numberOfBoxes * box[1]
            remainingSpace -= numberOfBoxes
            if (remainingSpace == 0) break
        }
        return loadedUnits
    }
}