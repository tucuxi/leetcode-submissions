func countPartitions(nums []int) int {
    allEven := true
    for _, num := range nums {
        allEven = allEven == (num % 2 == 0)
    }
    if allEven {
        return len(nums) - 1
    }
    return 0
}

/*
allEven | prefix even | suffix even | diff even
-----------------------------------------------
true    | true        | true        | true
true    | false       | false       | true
false   | true        | false       | false
false   | false       | true        | false
*/