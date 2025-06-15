type node struct {
    value int
    child [26]*node
}

type MapSum struct {
    root *node
}


func Constructor() MapSum {
    return MapSum{}
}


func (this *MapSum) Insert(key string, val int)  {

    var insert func(*node, string) *node
    insert = func(p *node, path string) *node {
        if p == nil {
            p = &node{}
        }
        if path == "" {
            p.value = val
            return p
        }
        ch := int(path[0] - 'a')
        p.child[ch] = insert(p.child[ch], path[1:])
        return p
    }

    this.root = insert(this.root, key)
}


func (this *MapSum) Sum(prefix string) int {
    
    var sum func(*node) int
    sum = func(p *node) int {
        if p == nil {
            return 0
        }
        res := p.value 
        for _, q := range p.child {
            res += sum(q)
        }
        return res
    }

    p := this.root
    for _, ch := range prefix {
        if p != nil {
            p = p.child[int(ch - 'a')]
        }
    }

    return sum(p)
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */