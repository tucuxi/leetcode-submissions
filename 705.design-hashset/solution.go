const size = 997

type MyHashSet struct {
    buckets [][]int
}


func Constructor() MyHashSet {
    return MyHashSet{make([][]int, size)}
}


func (this *MyHashSet) Add(key int)  {
    h := hash(key)
    for _, k := range this.buckets[h] {
        if k == key {
            return
        }
    }
    this.buckets[h] = append(this.buckets[h], key)
}


func (this *MyHashSet) Remove(key int)  {
    h := hash(key)
    b := this.buckets[h]
    for i := range b {
        if b[i] == key {
            this.buckets[h] = append(b[:i], b[i + 1:]...)
            return
        }
    }
}


func (this *MyHashSet) Contains(key int) bool {
    h := hash(key)
    for _, k := range this.buckets[h] {
        if k == key {
            return true
        }
    }
    return false
}


func hash(key int) int {
    return key % size
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */