/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    var sums []int64
    
    var traverse func(*TreeNode, int)
    traverse = func(root *TreeNode, level int) {
        if root == nil {
            return
        }
        if level == len(sums) {
            sums = append(sums, 0)
        }
        sums[level] += int64(root.Val)
        traverse(root.Left, level + 1)
        traverse(root.Right, level + 1)
    }
    
    traverse(root, 0)
    if k > len(sums) {
        return -1
    }
    return qselect(sums, 0, len(sums) - 1, len(sums) - k)
}

func qselect(a []int64, left, right, k int) int64 {
    for left < right {
        pi := partition(a, left, right, left + rand.Intn(right - left + 1))
        if pi == k {
            return a[k]
        }
        if k < pi {
            right = pi - 1
        } else {
            left = pi + 1
        }
    }
    return a[left]
}

func partition(a []int64, left, right, pivot int) int {
    pv := a[pivot]
    a[pivot], a[right] = a[right], a[pivot]
    st := left
    for i := left; i < right; i++ {
        if a[i] < pv {
            a[st], a[i] = a[i], a[st]
            st++
        }
    }
    a[st], a[right] = a[right], a[st]
    return st
}