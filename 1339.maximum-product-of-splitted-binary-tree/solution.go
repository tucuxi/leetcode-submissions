func maxProduct(root *TreeNode) int {
    sumRoot := sumTree(root)
    if sumRoot == nil {
        return 0
    }
    totalSum := sumRoot.Val
    var maxProd int64
    var recurse func(*TreeNode)
    recurse = func(root *TreeNode) {
        if root.Left != nil {
            if m := int64(root.Left.Val) * int64(totalSum - root.Left.Val); m > maxProd {
                maxProd = m
            }
            recurse(root.Left)
        }
        if root.Right != nil {
            if m := int64(root.Right.Val) * int64(totalSum - root.Right.Val); m > maxProd {
                maxProd = m
            }
            recurse(root.Right)
        }
    }

    recurse(sumRoot)
    return int(maxProd % 1_000_000_007)
}

func sumTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    sum := root.Val
    leftChild := sumTree(root.Left)
    rightChild := sumTree(root.Right)
    if leftChild != nil {
        sum += leftChild.Val
    }
    if rightChild != nil {
        sum += rightChild.Val
    }
    return &TreeNode{sum, leftChild, rightChild}
}