const (
    nbuckets = 113
)

type entry struct {
    key int
    val int
}

type MyHashMap struct {
    buckets [][]entry
}

func hash(key int) int {
    return key % nbuckets
}

func Constructor() MyHashMap {
    return MyHashMap{make([][]entry, nbuckets)}
}

func (this *MyHashMap) Put(key int, value int)  {
    h := hash(key)
    l := this.buckets[h]
    if l == nil {
        this.buckets[h] = []entry{entry{key, value}}
        return
    }
    for i := range l {
        if l[i].key == key {
            l[i].val = value
            return
        }
    }
    this.buckets[h] = append(l, entry{key, value})
}


func (this *MyHashMap) Get(key int) int {
    h := hash(key)
    l := this.buckets[h]
    for _, e := range l {
        if e.key == key {
            return e.val
        }
    }
    return -1
}

func (this *MyHashMap) Remove(key int)  {
    h := hash(key)
    l := this.buckets[h]
    for i := range l {
        if l[i].key == key {
            this.buckets[h] = append(l[:i], l[i+1:]...)
            return
        }
    }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */