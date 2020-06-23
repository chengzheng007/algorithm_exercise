package main

import (
	"fmt"
	"math"
)

func main() {
	w := 9
	weight := []int{2, 2, 4, 6, 3}
	maxw := zeroOnePackageDP(w, weight)
	fmt.Printf("zeroOnePackageDP maxw:%d\n", maxw)

	maxw = zeroOnePackageDP2(w, weight)
	fmt.Printf("zeroOnePackageDP2 maxw:%d\n", maxw)

	fmt.Println()

	dist := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	minDist := math.MaxInt64
	minDistDfs(len(dist), 0, 0, 0, &minDist, dist)
	fmt.Printf("minDistDfs:%v\n", minDist)

	minDist = minDistDP(dist)
	fmt.Printf("minDistDP:%v\n", minDist)

	fmt.Println()
	m := 12
	money := []int{1, 3, 5}
	fmt.Printf("find change money(%d) dfs num:%d\n", m, findChangeDfs(m, money))
	fmt.Printf("find change money(%d) dp num:%d\n\n", m, findChangeDP(m, money))

	fmt.Printf("combine change money(%d) max dfs num:%d\n", m, findChangeMaxDFS(money, m))
	fmt.Printf("combine change money(%d) max dp num:%d\n", m, findChangeMaxDP(money, m))

	a := "mitcmu"
	b := "mtacnu"
	fmt.Printf("lwst dfs min:%d\n", lwstDFS(a, b))
	fmt.Printf("lwst dp min:%d\n", lwstDP(a, b))
	fmt.Printf("lwst dp2 min:%d\n\n", lwstDP2(a, b))

	fmt.Printf("lcs dfs max:%d\n", lcsDFS(a, b))
	fmt.Printf("lcs dp max:%d\n", lcsDP(a, b))
	fmt.Printf("lcs dp2 max:%d\n", lcsDP(a, b))
	fmt.Printf("short common super sequence:%s\n\n", shortestCommonSupersequence(a, b))

	digit := []int{2, 2, 3}
	fmt.Printf("lis dfs max:%d\n", lisDFS(digit))
	fmt.Printf("lis dp max:%d\n", lisDP(digit))
	fmt.Printf("lis dfs2 len:%d\n\n", lengthOfLISDFS(digit, math.MinInt64, 0))

	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	n := 4
	fmt.Printf("cut rod dfs:%d\n", cutRodDFS(price, n, 0))
	fmt.Printf("cut rod dp:%d\n", cutRodDP(price, n))

	str := "bababvv"
	fmt.Printf("lps dfs len:%v\n", longestPalindromeDFS(str))
	fmt.Printf("lps dp len:%v\n", longestPalindromeDP(str))
}
