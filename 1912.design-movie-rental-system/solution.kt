class MovieRentingSystem(n: Int, entries: Array<IntArray>) {

    private data class Key(
        val shop: Int,
        val movie: Int,
    ) {
        constructor(arr: IntArray) : this(arr[0], arr[1])
    }

    private data class Entry(
        val shop: Int,
        val movie: Int,
        val price: Int,
    ) {
        constructor(arr: IntArray) : this(arr[0], arr[1], arr[2])
    }

    private val comparator = compareBy<Entry>({ it.price }, { it.shop }, { it.movie })

    private val movies = mutableMapOf<Key, Entry>()

    private val unrented = mutableMapOf<Int, TreeSet<Entry>>()

    private val rented = TreeSet<Entry>(comparator)

    init {
        entries.forEach { e ->
            val entry = Entry(e)
            movies[Key(e)] = entry
            unrented.getOrPut(entry.movie) { TreeSet<Entry>(comparator) }.add(entry)
        }
    }

    fun search(movie: Int): List<Int> {
        val shops = unrented[movie] ?: return emptyList()
        return shops.take(5).map { it.shop }
    }

    fun rent(shop: Int, movie: Int) {
        val key = Key(shop, movie)
        val entry = movies[key] ?: return
        unrented.getValue(movie).remove(entry)
        rented.add(entry)
    }

    fun drop(shop: Int, movie: Int) {
        val key = Key(shop, movie)
        val entry = movies[key] ?: return
        unrented.getValue(movie).add(entry)
        rented.remove(entry)
    }

    fun report(): List<List<Int>> {
        return rented.take(5).map { listOf(it.shop, it.movie) }
    }
}