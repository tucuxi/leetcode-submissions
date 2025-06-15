type vv struct {
    value    int
    version  int
}

type SnapshotArray struct {
    elems [][]vv
    snap  int
}


func Constructor(length int) SnapshotArray {
    return SnapshotArray{
        elems: make([][]vv, length),
        snap:  0,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    elem := this.elems[index]
    if len(elem) == 0 || elem[len(elem) - 1].version < this.snap {
        this.elems[index] = append(this.elems[index], vv{val, this.snap})
    } else {
        elem[len(elem) - 1].value = val
    }
}


func (this *SnapshotArray) Snap() int {
    id := this.snap
    this.snap++
    return id
}


func (this *SnapshotArray) Get(index int, snapId int) int {
    elem := this.elems[index]
    k := sort.Search(len(elem), func(i int) bool {
        return elem[i].version > snapId
    })
    if k == 0 {
        return 0
    }
    return elem[k - 1].value
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */