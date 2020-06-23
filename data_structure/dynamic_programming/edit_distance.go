package main

import (
	"fmt"
	"math"
)

// 编辑距离
// 莱文斯坦距离：刻画两个字符串之间的差异性
func lwstDFS(a, b string) int {
	min := math.MaxInt64
	lwstBacktrack(a, b, 0, 0, 0, &min)
	return min
}

func lwstBacktrack(a, b string, i, j, curr int, minDist *int) {
	if i == len(a) || j == len(b) {
		if i < len(a) {
			curr += len(a) - i
		}
		if j < len(b) {
			curr += len(b) - j
		}
		if curr < *minDist {
			*minDist = curr
		}
		return
	}
	if a[i] == b[j] {
		lwstBacktrack(a, b, i+1, j+1, curr, minDist)
	} else {
		lwstBacktrack(a, b, i+1, j, curr+1, minDist)   // 删除a[i]或者在b[j]前插入与a[i]相同的字符
		lwstBacktrack(a, b, i, j+1, curr+1, minDist)   // 删除b[j]或者在a[i]前插入与b[j]相同的字符
		lwstBacktrack(a, b, i+1, j+1, curr+1, minDist) // 将a[i]、b[j]复制为相同的字符
	}
}

/**
状态转移公式：
[i,j]这个状态从三个状态转移过来[i,j-1]、[i-1,j]、[i-1,j-1]（从递归也能看出）
设dp[i][j]为对比到a[i]和b[j]时的最小莱文斯坦距离
如果a[i] != b[j]，dp[i][j] = min(
		dp[i-1][j]+1, // 从(i-1,j)到(i,j)，i+1，j不变，说明删除了i或者b[j]前插入字符，所以距离+1
		dp[i][j-1]+1, // 操作同上
		dp[i-1][j-1]+1 //将a[i]替换为b[j]或相反
	)
如果a[i] == b[j]，dp[i][j] = min(
		dp[i-1][j-1],  // 由a[i-1]、b[j-1]对比到a[i][j]时，如果a[i]==b[j]，那么dp[i][j]==dp[i-1][j-1]
		dp[i-1][j]+1,
		dp[i][j-1]+1
	)

换个方式理解，(i,j)由(i-1,j)、(i,j-1)、(i-1,j-1)三个状态转移过来，从(i-1,j)或(i,j-1)到(i,j)
肯定要对字符做些操作，才能使一个下标前进，一个不前进，因此dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j-1]+1, "待定")
再对这里的待定分a[i]、b[j]等和不等的情况处理
*/
func lwstDP(a, b string) int {
	alen := len(a)
	blen := len(b)

	if alen == 0 {
		return blen
	} else if blen == 0 {
		return alen
	}

	dp := make([][]int, alen)
	for i := 0; i < alen; i++ {
		dp[i] = make([]int, blen)
	}

	// 初始化
	// 初始化dp第0行
	// 处理[a0]和[b0,b1....bn-1]的编辑距离
	// 注意一个序列长度是1(a0)，另一个是全长
	for j := 0; j < blen; j++ {
		if a[0] == b[j] {
			// 当b中任意一个字符跟a[0]相等时，距离值等于(j+1)-1=j
			// (j+1)是此时j下标为止字符串长度，a序列长度为1，那么len(b)-1
			// 个字符一定与a中的不同，因为a序列后面已经没有字符，所以长度就是len(b)-1
			dp[0][j] = j
		} else if j != 0 {
			// 若j不为0且b[j]不等于a[0]，而这里不存在[i-1][j]和[i-1][j-1]
			// 按照状态转移方程只能选dp[i][j-1]+1，即dp[0][j-1]+1
			// 另外因为存在第三种和第一种if的情况，所以这里需要在前一个基础上累加
			dp[0][j] = dp[0][j-1] + 1
		} else {
			// 特例，j==0且a[0]!=b[0]，编辑距离为1，可以实际验证
			// 如果j==0且a[0]==b[0]时，编辑距离为0，是第一种if的情况
			dp[0][j] = 1
		}
	}
	// 初始化第0列
	for i := 0; i < alen; i++ {
		if a[i] == b[0] {
			dp[i][0] = i
		} else if i != 0 {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = 1
		}
	}

	for i := 1; i < alen; i++ {
		for j := 1; j < blen; j++ {
			if a[i] == b[j] {
				dp[i][j] = minInt(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]) // 可以加dp[i-1][j-1]+1，但一定不是最小的
			} else {
				dp[i][j] = minInt(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	fmt.Println("lwst dp state:")
	for i := 0; i < alen; i++ {
		fmt.Printf("%v\n", dp[i])
	}

	return dp[alen-1][blen-1]
}

func minInt(args ...int) int {
	min := math.MaxInt64
	for _, num := range args {
		if num < min {
			min = num
		}
	}
	return min
}

func lwstDP2(word1, word2 string) int {
	size1 := len(word1)
	size2 := len(word2)
	if size1 == 0 {
		return size2
	}
	if size2 == 0 {
		return size1
	}
	dp := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		dp[i] = make([]int, size2+1)
	}

	// dp[i][j]表示word1[0:i)与word2[0:j)的距离值，注意末尾不包含i、j位置的索引

	// 初始化
	// dp[i][0]表示一个长度为i与一个空串的编辑距离，实际等于非空串本身长度
	// 注意下标从1开始，到本身长度值结束
	for i := 1; i <= size1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= size2; j++ {
		dp[0][j] = j
	}

	// 状态转移
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minInt(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[size1][size2]
}
