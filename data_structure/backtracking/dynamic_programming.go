package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// dp exercise
// https://www.zhihu.com/question/23995189
// 每个阶段只有一个状态->递推；
// 每个阶段的最优状态都是由上一个阶段的最优状态得到的->贪心；
// 每个阶段的最优状态是由之前所有阶段的状态的组合得到的->搜索；
// 每个阶段的最优状态可以从之前某个阶段的某个或某些状态直接得到而不管之前这个状态是如何得到的->动态规划。
//
// 每个阶段的最优状态可以从之前某个阶段的某个或某些状态直接得到
// 这个性质叫做最优子结构；

// 而不管之前这个状态是如何得到的
// 这个性质叫做无后效性。

// fibonacci number
// call: fibonacciNumDP(num)
func fibonacciNumDP(n int) int {
	// 假设n>=0
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	state := make([]int, n+1)

	// 初始化状态
	state[0] = 0
	state[1] = 1

	for i := 2; i <= n; i++ { // 状态转移
		// 转移方程：f(i) = f(i-1)+f(i-2)
		state[i] = state[i-1] + state[i-2]
	}
	return state[n]
}

func fibonacciNumBT(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	n2 := fibonacciNumBT(n - 2)
	n1 := fibonacciNumBT(n - 1)
	return n1 + n2
}

// 钱币找零
// 找回所选的金额所需要最少的个数
var (
	rmbs        = []int{1, 5, 11}
	currNeedNum = math.MaxInt32
)

func findChangeBT(i, cw, num int, rmbs []int, n int, w int) {
	if i == n || cw == w {
		//fmt.Printf("i(%d) == n || cw(%d) == w\n", i, cw)
		if num < currNeedNum {
			currNeedNum = num
		}
		return
	}
	for i := 0; i < n; i++ {
		// 如果当前选择的零钱加上已有零钱小于等于需找回的零钱，继续选择用当前零钱进行找回
		if cw+rmbs[i] <= w {
			//fmt.Printf("choose %d rmb:%d num:%d cw:%d cw+rmbs[i]:%d\n", i, rmbs[i], num+1, cw, cw+rmbs[i])
			findChangeBT(i, cw+rmbs[i], num+1, rmbs, n, w)
		}
	}
}

// 杨辉三角

// longest increasing substring/最长增长子序列
var (
	series = []int{1, 5, 3, 4, 6, 9, 7, 8}
)

func lisDP(list []int) int {
	size := len(list)
	if size <= 1 {
		return size
	}
	state := make([]int, size)
	// 初始化
	for i := range state {
		state[i] = 1
	}

	fmt.Printf("state[0]=%d\n", state[0])
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if list[j] < list[i] {
				state[i] = int(math.Max(float64(state[i]), float64(state[j]+1)))
			}
		}
		fmt.Printf("state[%d]=%d\n", i, state[i])
	}
	max := state[0]
	for i := 1; i < size; i++ {
		if max < state[i] {
			max = state[i]
		}
	}
	return max
}

// 全排列
func permutation(list []int) [][]int {
	res := make([][]int, 0)
	dep := 0
	permutate(list, len(list), 0, &res, dep)
	return res
}

func permutate(list []int, n, start int, res *[][]int, dep int) {
	if start == n {
		// 这里必须临时复制一份，因为各个函数会对nums造成影响
		// 递归返回后，将数据换回来，导致res中
		// 的排列都是换回来的结果
		temp := make([]int, n)
		copy(temp, list)
		*res = append(*res, temp)
		return
	}
	// i从初始的start值开始，与start位置的元素交换，如果start位置本身递归完成回来后
	// i+1再和start位置交换继续打印
	for i := start; i < n; i++ {
		fmt.Printf("dep:%d, swap i(%d) and start(%d)\n", dep, i, start)
		list[i], list[start] = list[start], list[i]
		permutate(list, n, start+1, res, dep+1)
		fmt.Printf("dep:%d, revoke swap i(%d) and start(%d)\n", dep, i, start)
		list[i], list[start] = list[start], list[i]
	}
}

func permutation2(list []int) [][]int {
	var res [][]int
	var path []int
	visited := make([]bool, len(list))
	res = permuteBacktrack(list, path, visited, res)
	return res
}

// 这里的思想比交换更简单，但额外消耗了空间，采用一个map，记录当前已经选了哪些数
// 当循环中检测到当前数已经被选时就跳过，当检测到待选数track总长度与总数长度相等时，加入到结果集中
// 注意递归回退时，需将路径数据回归原位
func permuteBacktrack(nums []int, path []int, visited []bool, res [][]int) [][]int {
	if len(path) == len(nums) {
		arr := make([]int, len(nums))
		copy(arr, path)
		res = append(res, arr)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			// 选择
			path = append(path, nums[i])
			visited[i] = true
			res = permuteBacktrack(nums, path, visited, res)
			// 撤销选择
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	return res
}

// pos[i]表示第i行摆放的列数
func nQueen(pos []int, n, row int, count *int) {
	if len(pos) != n {
		return
	}
	if row == n {
		*count++
		printQueen(pos, n)
		return
	}
	for column := 0; column < n; column++ {
		if isOK(pos, n, row, column) {
			pos[row] = column
			nQueen(pos, n, row+1, count)
			pos[row] = 0
		}
	}
}

// 只需要找到一种放置方法即可
func nQueen2(pos []int, n, row int) bool {
	if len(pos) != n {
		return false
	}
	if row == n {
		return true
	}
	for col := 0; col < n; col++ {
		if isOK(pos, n, row, col) {
			pos[row] = col
			// 注意这里找到也需要return
			if nQueen2(pos, n, row+1) {
				return true
			}
			pos[row] = 0
		}
	}
	return false
}

func isOK(pos []int, n, row, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for i := row - 1; i >= 0; i-- {
		// 检测上一行该列处是否有位置
		if pos[i] == column {
			return false
		}
		// 检测左上角
		if leftUp >= 0 && pos[i] == leftUp {
			return false
		}
		if rightUp < n && pos[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}

func printQueen(pos []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if pos[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// 0-1背包问题
// 每件物品不可分割，总重量不能超过w，求出最多装入重量是多少
func zeroOnePackage(w int, weight []int, cw, i int, maxw *int) {
	if cw == w || i == len(weight) {
		if cw > *maxw {
			*maxw = cw
		}
		return
	}
	// 先不装第i件物品
	zeroOnePackage(w, weight, cw, i+1, maxw)
	// 装入第i件，注意这里做了剪枝，先校验是否已超过w，防止过多无意义递归试探
	if cw+weight[i] <= w {
		zeroOnePackage(w, weight, cw+weight[i], i+1, maxw)
	}
}

func zeroOnePackage2(w int, weight []int, n, k, cw int, maxw *int) {
	if cw == w || k == n {
		if cw > *maxw {
			*maxw = cw
		}
		return
	}
	for i := k; i < n; i++ {
		if cw+weight[i] <= w {
			zeroOnePackage2(w, weight, n, i+1, cw+weight[i], maxw)
		}
	}
}

func rmatch(text string, tlen int, reg string, rlen int, matched *bool, ti, pj int) {
	// 已匹配
	if *matched {
		return
	}
	if pj == rlen {
		if ti == tlen {
			*matched = true
		}
		return
	}
	if reg[pj] == '*' {
		for k := 0; k < tlen-ti; k++ {
			rmatch(text, tlen, reg, rlen, matched, ti+k, pj+1)
		}
	} else if reg[pj] == '?' {
		rmatch(text, tlen, reg, rlen, matched, ti, pj+1)
		rmatch(text, tlen, reg, rlen, matched, ti+1, pj+1)
	} else if ti < tlen && text[ti] == reg[pj] {
		rmatch(text, tlen, reg, rlen, matched, ti+1, pj+1)
	}
}

func subset(nums []int) [][]int {
	res := make([][]int, 0)
	curSet := make([]int, 0)
	count := 0
	subsetBacktrack(nums, len(nums), 0, curSet, &res, &count)
	fmt.Printf("count:%v\n", count)
	return res
}

func subsetBacktrack(nums []int, n, k int, curSet []int, res *[][]int, count *int) {
	*res = append(*res, curSet)
	*count++
	for i := k; i < n; i++ {
		curSet2 := make([]int, len(curSet)+1)
		copy(curSet2, curSet)
		curSet2[len(curSet)] = nums[i]
		fmt.Printf("i:%d, k:%d, curSet2:%v\n", i, k, curSet2)
		subsetBacktrack(nums, n, i+1, curSet2, res, count)
	}
}

func subset2(nums []int) [][]int {
	res := make([][]int, 0)
	curSet := make([]int, 0)
	subsetBacktrack2(nums, len(nums), 0, curSet, &res)
	return res
}

func subsetBacktrack2(nums []int, n, i int, curSet []int, res *[][]int) {
	if i == n {
		*res = append(*res, curSet)
		return
	}
	// 不选第i个元素
	subsetBacktrack2(nums, n, i+1, curSet, res)

	// 选第i个元素
	curSet2 := make([]int, len(curSet)+1)
	copy(curSet2, curSet)
	curSet2[len(curSet)] = nums[i]
	subsetBacktrack2(nums, n, i+1, curSet2, res)
}

// 当数组有重复元素时，集合中需剔除重复，如[1,2,2]
// 生成的集合中不能同时有两个[2]或两个[1,2]
func susbetDup(nums []int) [][]int {
	// 先从小到大排序
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int, 0)
	curSet := make([]int, 0)
	subsetDupBacktrack(nums, len(nums), 0, curSet, &res)
	return res
}

func subsetDupBacktrack(nums []int, n, k int, curSet []int, res *[][]int) {
	*res = append(*res, curSet)
	for i := k; i < n; i++ {
		// Note:为什么判断的是i == i-1，而不是i == i+1?
		// 回溯是从小的下标到大的下标，前面的元素先参与到集合子元素的组建
		// 所以后面的数如果与前面的相同，肯定已经参与过组建了
		if i > k && nums[i] == nums[i-1] {
			continue
		}
		curSet2 := make([]int, len(curSet)+1)
		copy(curSet2, curSet)
		curSet2[len(curSet)] = nums[i]
		subsetDupBacktrack(nums, n, i+1, curSet2, res)
	}
}

func rmatch2(s, p string, si, pj int) bool {
	if pj == len(p) {
		fmt.Printf("if pj:%d, si:%d\n", pj, si)
		if si == len(s) {
			return true
		}
		return false
	}
	if p[pj] == '*' {
		fmt.Printf("match *, si:%d, pj:%d, len(s)-si:%d\n", si, pj, len(s)-si)
		for k := 0; k <= len(s)-si; k++ {
			if rmatch2(s, p, si+k, pj+1) {
				return true
			}
		}
	} else if p[pj] == '.' {
		if rmatch2(s, p, si+1, pj+1) {
			return true
		}
	} else if si < len(s) && s[si] == p[pj] {
		if rmatch2(s, p, si+1, pj+1) {
			return true
		}
	}
	return false
}

func letterCasePermutation(S string) []string {
	var res []string
	for i := 0; i < len(S); i++ {
		ret := isLetter(S[i])
		size := len(res)
		if ret == 1 { // 小写字母
			var temp []string
			if size == 0 {
				temp = make([]string, 2)
				temp[0] = string(S[i])
				temp[1] = string(S[i] - 32)
			} else {
				temp = make([]string, 2*size)
				for j := 0; j < size; j++ {
					// fmt.Printf("temp[%d]:%v\n", j, temp[j])
					temp[j] = res[j] + string(S[i])
					temp[j+size] = res[j] + string(S[i]-32)
					// fmt.Printf("after temp[%d]:%v, temp[%d]:%v\n", j, temp[j], j+size, temp[j+size])
				}
			}
			res = temp
			// fmt.Printf("lower res:%v\n", res)
		} else if ret == 2 { // 大写字母
			var temp []string
			if size == 0 {
				temp = make([]string, 2)
				temp[0] = string(S[i])
				temp[1] = string(S[i] + 32)
			} else {
				temp = make([]string, 2*size)
				for j := 0; j < size; j++ {
					temp[j] = res[j] + string(S[i])
					temp[j+size] = res[j] + string(S[i]+32)
				}
			}
			res = temp
			// fmt.Printf("upper res:%v\n", res)
		} else { // 非字母
			for j := range res {
				res[j] = res[j] + string(S[i])
			}
			// fmt.Printf("normal res:%v\n", res)
		}
	}
	return res
}

// 0不是字母 1小写字母 2大写字母
func isLetter(s byte) int {
	if s >= 'a' && s <= 'z' {
		return 1
	} else if s >= 'A' && s <= 'Z' {
		return 2
	}
	return 0
}

func restoreIpAddresses(s string) []string {
	var res []string
	var path []string
	backtrackIPAddress(s, len(s), 0, path, &res)
	return res
}

func backtrackIPAddress(s string, n, index int, path []string, res *[]string) {
	if index == 4 {
		if s == "" {
			*res = append(*res, strings.Join(path, "."))
		}
		return
	}
	for i := 1; i < 4; i++ {
		if i > len(s) {
			return
		}
		if isValid(s[:i]) {
			path = append(path, s[:i])
			backtrackIPAddress(s[i:], n, index+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isValid(str string) bool {
	if len(str) == 0 || len(str) > 3 {
		return false
	}
	if str[0] == '0' && len(str) > 1 {
		return false
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}

// leetcode351 Android Unlock Patterns-安卓解锁手势
// 在安卓手机的1...9的数字键盘上，用若干个点连接起来组成一种解锁手势，若干个点的长度位于[m,n]之间，求有多少种解法
// 手势要求：1.每个手势解锁密码至少包含m个数，至多包含n个数
// 2.手势中所有的数都不相同
// 3.两个顺序上连续的数（不是大小连续，是连接顺序上的先后）中间经过的点必须是之前已经选择连接过的
// 4.数的使用顺序很重要
// 要求解读：1.手势密码，也就是解锁需要的数字，至少m至多n个；2.不能重复连接同一个数两次；
// 3.一个数连接到另一个数时，中间会经过其他的数，该数必须是之前已经选择过（已经连接的），比如1连到3
// 中间经过2，那么2必须是之前已经使用的了；4.即便数字相同连接顺序不同也算作不同的连接方式
// 分析：第1、2点表示手势中的数必须都不同，因此可以设置一个visit数组，标记该数字是否被使用过
// 第3点规定两个连续的点中间经过的数必须已被使用，哪些两个数连接时中间需要经过数字，1、3中间有2，4、6
// 中间有5，总结下四个对角线(1-9,9-1算两个)和四条直线(2-8,8-2算两条)都经过5，其余是1-3、1-7
// 7-9等，可以使用二维数组预先记录下这些信息，等到需要连接下一个数字时，检测他们之间是否会经过其他数字
// 由于1、3、7、9是对称的，求出1的个数乘4即可，2、4、6、8求出2的个数乘4即可，剩余5求一次个数，整个相加
// 如果经过是否已被使用(在visited中查询状态)等等
// 大致思路：深度优先的回溯算法
func numberOfPatterns(m, n int) int {
	// 总共9个数，以数字直接访问下标，true表示已访问
	visited := make([]bool, 10)
	jump := make([][]int, 10)
	for i := 1; i <= 9; i++ {
		jump[i] = make([]int, 10)
	}

	// jump记录两个键盘数字之间是否经过其他数，不经过则值为0
	jump[1][3], jump[3][1] = 2, 2
	jump[3][9], jump[9][3] = 6, 6
	jump[9][7], jump[7][9] = 8, 8
	jump[7][1], jump[1][7] = 4, 4
	jump[1][9], jump[9][1] = 5, 5
	jump[2][8], jump[8][2] = 5, 5
	jump[3][7], jump[7][3] = 5, 5
	jump[4][6], jump[6][4] = 5, 5

	res := getPatternCnt(1, 1, 0, m, n, jump, visited) * 4
	res += getPatternCnt(2, 1, 0, m, n, jump, visited) * 4
	res += getPatternCnt(5, 1, 0, m, n, jump, visited)
	return res
}

func getPatternCnt(start, length, cnt, m, n int, jump [][]int, visited []bool) int {
	if length >= m {
		cnt++
	}
	length++
	// 超过了最大长度（最多点数）
	if length > n {
		return cnt
	}
	visited[start] = true
	// 为什么从1开始循环？即便是9也可以连接到1
	for next := 1; next <= 9; next++ {
		passby := jump[start][next]
		// 下一个数没访问过，且经过下一个数时不经过其他数，或者经过其他数但该数已被访问
		if !visited[next] && (passby == 0 || visited[passby]) {
			// 不是+=，即便这里有递归，也是上面自增调用后递归回退
			cnt = getPatternCnt(next, length, m, n, cnt, jump, visited)
		}
	}
	visited[start] = false
	return cnt
}
