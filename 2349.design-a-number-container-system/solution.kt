class NumberContainers() {

    private val indexToNumber = mutableMapOf<Int, Int>()

    private val numberToIndex = mutableMapOf<Int, SortedSet<Int>>()

    fun change(index: Int, number: Int) {
        indexToNumber[index]?.let { numberToIndex[it]?.remove(index) }
        indexToNumber[index] = number
        numberToIndex.getOrPut(number) { sortedSetOf() }.add(index)
    }

    fun find(number: Int): Int {
        return numberToIndex[number]?.firstOrNull() ?: -1
    }

}