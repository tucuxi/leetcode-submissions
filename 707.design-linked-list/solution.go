type MyLinkedList struct {
    sentinel *myEntry
    length int
}

type myEntry struct {
    value int
    previous, next *myEntry
}

func Constructor() MyLinkedList {
    sentinel := &myEntry{}
    sentinel.previous = sentinel
    sentinel.next = sentinel
    return MyLinkedList{sentinel, 0}
}

func (this *MyLinkedList) Get(index int) int {
    if index >= this.length {
        return -1
    }
    node := advance(this.sentinel.next, index)
    return node.value
}

func (this *MyLinkedList) addLeft(currentNode *myEntry, value int) {
    newNode := &myEntry{value, currentNode.previous, currentNode}
    currentNode.previous.next = newNode
    currentNode.previous = newNode
    this.length++
}

func (this *MyLinkedList) addRight(currentNode *myEntry, value int) {
    newNode := &myEntry{value, currentNode, currentNode.next}
    currentNode.next.previous = newNode
    currentNode.next = newNode
    this.length++
}

func (this *MyLinkedList) AddAtHead(val int)  {
    this.addRight(this.sentinel, val)
}

func (this *MyLinkedList) AddAtTail(val int)  {
    this.addRight(this.sentinel.previous, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    switch {
        case index < this.length:
            this.addLeft(advance(this.sentinel.next, index), val)
        case index == this.length:
            this.addRight(this.sentinel.previous, val)
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.length {
        return
    }
    node := advance(this.sentinel.next, index)
    node.previous.next = node.next
    node.next.previous = node.previous
    this.length--
}

func advance(start *myEntry, steps int) *myEntry {
    p := start
    for steps > 0 {
        p = p.next
        steps--
    }
    return p
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */