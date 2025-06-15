type node struct {
    count int
    keys map[string]bool
    previous, next *node
}

func newNode(count int) *node {
    return &node{
        count: count,
        keys: make(map[string]bool),
    }
}

type AllOne struct {
    nodes map[string]*node
    head, tail *node
}

func Constructor() AllOne {
    head := newNode(0)
    tail := newNode(0)
    head.next = tail
    tail.previous = head
    return AllOne{
        nodes: make(map[string]*node),
        head: head,
        tail: tail,
    }
}

func (this *AllOne) Inc(key string)  {
    node, exists := this.nodes[key]
    if exists {
        delete(node.keys, key)
        count := node.count
        next := node.next
        if next.count != count + 1 {
            newNode := newNode(count + 1)
            newNode.keys[key] = true
            newNode.previous = node
            newNode.next = next
            node.next = newNode
            next.previous = newNode
            this.nodes[key] = newNode
        } else {
            next.keys[key] = true
            this.nodes[key] = next
        }
        if len(node.keys) == 0 {
            node.previous.next, node.next.previous = node.next, node.previous
            node.next, node.previous = nil, nil
        }
    } else {
        first := this.head.next
        if first.count != 1 {
            newNode := newNode(1)
            newNode.keys[key] = true
            newNode.previous = this.head
            newNode.next = first
            this.head.next = newNode
            first.previous = newNode
            this.nodes[key] = newNode
        } else {
            first.keys[key] = true
            this.nodes[key] = first
        }
    }
}

func (this *AllOne) Dec(key string)  {
    node, exists := this.nodes[key]
    if !exists {
        return
    }
    delete(node.keys, key)
    delete(this.nodes, key)
    if count := node.count; count > 1 {
        previous := node.previous
        if previous == this.head || previous.count != count - 1 {
            newNode := newNode(count - 1)
            newNode.keys[key] = true
            newNode.previous = previous
            newNode.next = node
            previous.next = newNode
            node.previous = newNode
            this.nodes[key] = newNode
        } else {
            previous.keys[key] = true
            this.nodes[key] = previous
        }
    }
    if len(node.keys) == 0 {
        node.previous.next, node.next.previous = node.next, node.previous
        node.next, node.previous = nil, nil
    }
}

func (this *AllOne) GetMaxKey() string {
    for key := range this.tail.previous.keys {
        return key
    }
    return ""
}

func (this *AllOne) GetMinKey() string {
    for key := range this.head.next.keys {
        return key
    }
    return ""
}

func (this *AllOne) String() string {
    var sb strings.Builder
    for p := this.head; p != nil; p = p.next {
        s := fmt.Sprintf("node{count: %d, previous: %p, next: %p}", p.count, p.previous, p.next)
        sb.WriteString(s + ",")
    }
    return sb.String()
}