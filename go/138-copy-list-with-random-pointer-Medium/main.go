/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    m := make(map[*Node]*Node)
    for node := head; node != nil; node = node.Next {
        m[node] = new(Node)
    }
    for node := head; node != nil; node = node.Next {
        newNode := m[node]
        newNode.Val = node.Val
        newNode.Next = m[node.Next]
        newNode.Random = m[node.Random]
    }
    return m[head]
}