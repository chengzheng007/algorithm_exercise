func canCompleteCircuit(gas []int, cost []int) int {
	start := len(gas) - 1
	end := 0
	sum := gas[start] - cost[start]
	for start > end {
		if sum >= 0 {
			sum += gas[end] - cost[end]
			end++
		} else {
			start--
			sum += gas[start] - cost[start]
		}
	}

	if sum >= 0 {
		return start
	}

	return -1
}

// O(n^2)
// func canCompleteCircuit(gas []int, cost []int) int {
//     n := len(gas)
//     for i := 0; i < n; i++ {
//         curr := gas[i]
//         j := (i+1)%n
//         for ; j != i; j = (j+1)%n {
//             if curr - cost[(j-1+n)%n] < 0 {
//                 // 无法到达下一站
//                 break
//             }
//             curr = curr-cost[(j-1+n)%n]+gas[j]
//         }
//         // 验证最后一站是否能到达
//         if j == i && curr-cost[(j-1+n)%n] >= 0 {
//             return i
//         }
//     }
//     return -1
// }