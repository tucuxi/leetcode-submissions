// Kotlin Iterator reference:
// https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.collections/-iterator/

class PeekingIterator(val iterator: Iterator<Int>): Iterator<Int> {
    
    private var next = 0
    
    init { next() }
    
    fun peek(): Int {
    	return next ?: throw NoSuchElementException() 
    }
    
    override fun next(): Int {
        return next.also {
            next = if (iterator.hasNext()) iterator.next() else 0
        }
    }
    
    override fun hasNext(): Boolean {
        return next != 0
    }
}

/**
 * Your PeekingIterator object will be instantiated and called as such:
 * var obj = PeekingIterator(arr)
 * var param_1 = obj.next()
 * var param_2 = obj.peek()
 * var param_3 = obj.hasNext()
 */