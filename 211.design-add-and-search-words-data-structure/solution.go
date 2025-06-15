type trieNode struct {
    children map[rune]*trieNode
    terminal bool
}

type WordDictionary struct {
    trieRoot *trieNode
}


func Constructor() WordDictionary {
    nn := new(trieNode)
    nn.children = make(map[rune]*trieNode)
    return WordDictionary{nn}
}


func (this *WordDictionary) AddWord(word string) {
    tn := this.trieRoot
    for _, r := range word {
        if tn.children[r] == nil {
            nn := new(trieNode)
            nn.children = make(map[rune]*trieNode)
            tn.children[r] = nn
        }
        tn = tn.children[r]
    }
    tn.terminal = true
}


func (this *WordDictionary) Search(word string) bool {
    tns := []*trieNode{this.trieRoot}
    for _, r := range word {
        k := len(tns)
        for _, tn := range tns {
            if r == '.' {
                for _, c := range tn.children {
                    tns = append(tns, c)
                }
            } else {
                if tn.children[r] != nil {
                    tns = append(tns, tn.children[r])
                }
            }
        }
        tns = tns[k:]
    }
    for _, tn := range tns {
        if tn.terminal {
            return true
        }
    }
    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */