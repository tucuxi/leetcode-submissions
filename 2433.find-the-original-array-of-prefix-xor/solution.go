func findArray(pref []int) []int {
    arr := make([]int, len(pref))
    a := 0
    for i := range pref {
        arr[i] = pref[i] ^ a
        a ^= arr[i]
    } 
    return arr
}

/*

pref[0] = arr[0]
arr[0] = pref[0] ^ 0

pref[1] = arr[0] ^ arr[1]
arr[1] = pref[1] ^ arr[0]

pref[2] = arr[0] ^ arr[1] ^ arr[2]
arr[2] = pref[2] ^ arr[0] ^ arr[1]

*/