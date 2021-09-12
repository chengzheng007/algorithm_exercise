func canCompleteCircuit(gas []int, cost []int) int {
    // O(N)的解法是，设置两个变量，sum判断当前的指针的有效性
    // total 则判断整个数组是否有解
    // 有就返回通过 sum 得到的下标，没有则返回 -1
    // total为所有的gas相加，减去所有的cost，只有这个差值大于等于0
    // 才可能存在解
    var (
        total, sum int
        start int
    )
    for i := 0; i < len(gas); i++ {
        sum = sum + gas[i] - cost[i]
        total += gas[i] - cost[i]
        if sum < 0 {
            start = i+1
            sum = 0
        }
    }
    
    if total >= 0 {
        return start
    }
    
    return -1
}

// O(n^2)
func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    for start := 0; start < n; start++ {
        // 当前储存的油气，如果到不了下一站，直接从下一站开始检测起点
        store := gas[start] - cost[start]
        if store < 0 {
            continue
        }
        station := (start+1)%n
        for station != start {
            // 计算到station下一站的油气余量
            store = store + gas[station] - cost[station]
            if store < 0 {
                break
            }
            station = (station+1)%n
        }
        // 能到达最开始的站点
        if station == start {
            return start
        }
    }
    return -1
}
