class EventManager(events: Array<IntArray>) {

    data class Event(val eventId: Int, val priority: Int) 

    private val map = mutableMapOf<Int, Event>()
    private val pq = PriorityQueue<Event>(compareBy({ -it.priority }, { it.eventId } )) 

    init {
        events.forEach { (eventId, priority) ->
            val event = Event(eventId, priority)
            map[eventId] = event
            pq.offer(event)
        }
    }

    fun updatePriority(eventId: Int, newPriority: Int) {
        val event = Event(eventId, newPriority)
        map[eventId] = event
        pq.add(event)
    }

    fun pollHighest(): Int {
        while (pq.isNotEmpty()) {
            val event = pq.poll()
            if (map[event.eventId] == event) {
                return event.eventId.also { map.remove(it) }
            }
        }
        return -1
    }
}

/**
 * Your EventManager object will be instantiated and called as such:
 * var obj = EventManager(events)
 * obj.updatePriority(eventId,newPriority)
 * var param_2 = obj.pollHighest()
 */