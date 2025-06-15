/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    n := mountainArr.length()
    p := sort.Search(n - 1, func(i int) bool {
        return mountainArr.get(i) > mountainArr.get(i + 1)
    })
    u := sort.Search(p, func(i int) bool {
        return mountainArr.get(i) >= target
    })
    if u < p && mountainArr.get(u) == target {
        return u
    }
    v := p + sort.Search(n - p, func(i int) bool {
        return mountainArr.get(p + i) <= target
    })
    if v < n && mountainArr.get(v) == target {
        return v
    }
    return -1
}