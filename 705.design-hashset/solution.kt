class MyHashSet() {

    private val table = Array<MutableList<Int>?>(997) { null }
    
    fun add(key: Int) {
        val h = key.hashCode() % table.size
        val l = table[h] ?: mutableListOf<Int>()
        if (key !in l) {
            l.add(key)
            table[h] = l
        }
    }

    fun remove(key: Int) {
        val h = key.hashCode() % table.size
        table[h]?.remove(key)        
    }

    fun contains(key: Int): Boolean {
        val h = key.hashCode() % table.size
        val l = table[h]
        return l?.contains(key) ?: false
    }

}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * var obj = MyHashSet()
 * obj.add(key)
 * obj.remove(key)
 * var param_3 = obj.contains(key)
 */