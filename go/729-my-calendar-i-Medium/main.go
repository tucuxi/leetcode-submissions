type MyCalendar struct {
    root *node 
}


func Constructor() MyCalendar {
    return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    if search(this.root, start, end) != nil {
        return false
    }
    if this.root == nil {
        this.root = &node{
            start: start,
            end: end,
            max: end,
        }
    } else {
        add(this.root, start, end)
    }
    return true
}


func add(root *node, start, end int) {
    prev, curr := root, root
    for curr != nil {
        prev = curr
        if end > curr.max {
            curr.max = end
        }
        if start < curr.start {
            curr = curr.left
        } else {
            curr = curr.right
        }
    }
    curr = &node{
        start: start,
        end: end,
        max: end,
    }   
    if start < prev.start {
        prev.left = curr
    } else {
        prev.right = curr
    }
}


func search(root *node, start, end int) *node {
    x := root
    for x != nil && (x.start >= end || x.end <= start) {
        if x.left != nil && x.left.max >= start {
            x = x.left
        } else {
            x = x.right
        }
    }
    return x
}


type node struct {
    start, end, max int
    left, right *node
}