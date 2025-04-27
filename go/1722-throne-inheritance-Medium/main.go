type ThroneInheritance struct {
    root *node
    dict map[string]*node
}


type node struct {
    name string
    children []*node
    isAlive bool
}


func Constructor(kingName string) ThroneInheritance {
    king := &node{name: kingName, isAlive: true}
    dict := map[string]*node{kingName: king}
    return ThroneInheritance{king, dict}
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {
    child := &node{name: childName, isAlive: true}
    this.dict[childName] = child
    parent := this.dict[parentName]
    parent.children = append(parent.children, child)
}


func (this *ThroneInheritance) Death(name string)  {
    this.dict[name].isAlive = false
}


func (this *ThroneInheritance) GetInheritanceOrder() []string {
    return preorder(this.root)
}


func preorder(p *node) []string {
    var res []string

    if p != nil {
        if p.isAlive {
            res = append(res, p.name)
        }
        for _, child := range p.children {
            res = append(res, preorder(child)...)
        }
    }
    return res
}


/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */