type node [10]*node

type trie struct {
    root node
}

func (t *trie) insert(a int) {
    p := &t.root
    for _, r := range strconv.Itoa(a) {
        n := int(r - '0')
        if p[n] == nil {
            p[n] = new(node)
            p = p[n]
        } else {
            p = p[n]
        }
    }
}

func (t *trie) prefix(a int) int {
    p := &t.root
    l := 0
    for _, r := range strconv.Itoa(a) {
        n := int(r - '0')
        p = p[n]
        if p == nil {
            break 
        }
        l++
    }
    return l
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
    t := new(trie)
    for _, a := range arr2 {
        t.insert(a)
    }
    res := 0
    for _, a := range arr1 {
        res = max(res, t.prefix(a))
    }
    return res
}