/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var res bytes.Buffer

    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            res.WriteByte(',')
            return
        }
        res.WriteString(strconv.Itoa(node.Val))
        res.WriteByte(',')
        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)
    return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    tokens := strings.Split(data, ",")
    
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if len(tokens) == 0 {
            return nil
        }
        token := tokens[0]
        tokens = tokens[1:]
        if token == "" {
            return nil
        }
        val, _ := strconv.Atoi(token)
        return &TreeNode{val, dfs(), dfs()}
    }

    return dfs()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */