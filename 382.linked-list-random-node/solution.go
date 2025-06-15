/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    head *ListNode
}


func Constructor(head *ListNode) Solution {
    return Solution{head}
}


func (this *Solution) GetRandom() int {
    res := 0
    n := 0
    for node := this.head; node != nil; node = node.Next {
        n++
        if rand.Intn(n) == 0 {
            res = node.Val
        }
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */