type LRUCache struct {
    data map[int]*entry
    capacity int
    sentinel *entry
}

type entry struct {
    key, value int
    previous, next *entry
}

func Constructor(capacity int) LRUCache {
    data := make(map[int]*entry, capacity)
    sentinel := &entry{}
    sentinel.previous = sentinel
    sentinel.next = sentinel
    return LRUCache{data, capacity, sentinel}
}

func (this *LRUCache) Get(key int) int {
    existing := this.data[key]
    if existing == nil {
        return -1
    }
    unlink(existing)
    link(existing, this.sentinel)
    return existing.value
}

func (this *LRUCache) Put(key int, value int)  {
    existing := this.data[key]
    switch {
        case existing == nil && len(this.data) == this.capacity:
            victim := this.sentinel.previous
            unlink(victim)
            delete(this.data, victim.key)
            victim.key = key
            victim.value = value
            this.add(victim)
        case existing == nil:
            newEntry := &entry{key: key, value: value}
            this.add(newEntry)
        default:
            existing.value = value
            unlink(existing)
            link(existing, this.sentinel)
    }
}

func (this *LRUCache) add(newEntry *entry) {
    link(newEntry, this.sentinel)
    this.data[newEntry.key] = newEntry
}

func link(e *entry, current *entry) {
    e.next = current.next
    e.previous = current
    current.next.previous = e
    current.next = e
}

func unlink(e *entry) {
    e.previous.next = e.next
    e.next.previous = e.previous
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */