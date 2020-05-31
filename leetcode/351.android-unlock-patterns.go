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