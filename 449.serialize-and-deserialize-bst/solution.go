/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    return fmt.Sprintf("%d:%s:%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    node, _ := deserialize(data)
    return node
}

func deserialize(data string) (*TreeNode, string) {
    if data == "" {
        return nil, ""
    }
    if data[0] == ':' {
        return nil, data[1:]
    }

    v := 0
    s := data
    
    for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
        v = 10*v + int(s[0] - '0')
        s = s[1:]
    }
    if len(s) > 0 {
        s = s[1:]
    }

    var l, r *TreeNode
    l, s = deserialize(s)
    r, s = deserialize(s)

    return &TreeNode{v, l, r}, s
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */