type timedValue struct {
    timestamp int
    value string
}

type TimeMap struct {
    entries map[string][]timedValue
}


func Constructor() TimeMap {    
    return TimeMap{map[string][]timedValue{}}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.entries[key] = append(this.entries[key], timedValue{timestamp, value})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    values := this.entries[key]
    if values == nil {
        return ""
    }
    greater := sort.Search(len(values), func(i int) bool {
        return values[i].timestamp > timestamp
    })
    if greater == 0 {
        return ""
    }
    return values[greater - 1].value
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */