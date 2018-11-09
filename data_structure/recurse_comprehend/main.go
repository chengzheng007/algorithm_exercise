package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("sum(%d):%d recursiveSum(%d):%d\n", n, sum(100), n, recursiveSum(100))

	n = 13
	fmt.Printf("recursie == %d step(s), there's %d ways\n", n, recursiveStep(n))
	fmt.Printf("non rec == %d step(s), there's %d ways\n", n, step(n))
}


// 找出递归规律，假设第n-1步成功了，第n步只需做什么，然后跟第n-1步建立联系
// 递归终止条件，一个还是多个
// 对于递归来说，正向是递，不断调用下一个函数将参数传递进去，回来是归，上一次计算结果退回
// 递归可以用递推表示，如 f(n) = f(n-1) + 1, f(1) = 1



// 求1-n的自然数的和，递推公式是f(n) = n + f(n-1)，要假设f(n-1)表示的是1~n-1的自然数的和，
// 那么1~n和就是在f(n-1)的基础上加n，不要非得一步步还原递归的每一步，容易陷进去

// 递推（非递归）
func sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
// 递归
func recursiveSum(n int) int {
	if n == 1 {
		return 1
	}
	return  n+recursiveSum(n-1)
}

// 上楼梯问题，每一步可以上一个台阶或者两个台阶，上n个台阶有多少种走法
// 思考，假设只有1个台阶，只有一种走法，假设只有2个台阶，就有两种走法：一步一步上或一次走两个
// 也就是最终多少种走法，依赖第一步怎么走，第一次走1步，n个台阶走法为1加上剩下n-1步台阶的走法f(n-1)
// 第一次走2步，n个台阶总走法等于1加上剩下n-2步台阶的走法f(n-2)
// 就是说递推公式是：f(n)=f(n-1)+f(n-2)，递归的终止条件是：f(1)=1,f(2)=2
// 假设终止条件只有f(1)=1，那么f(2)=f(1)+f(0),f(0)=1有点奇怪，0个台阶有一种走法
// 故可以将f(1)和f(2)作为终止条件

func recursiveStep(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return recursiveStep(n-1)+recursiveStep(n-2)
}

// 如何将上述算法转成非递归？需要顺着循环递推走
func step(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	totalSteps := 0
	prePre := 1  // 一个台阶走法
	pre := 2     // 两个个台阶走法
	for i := 3; i <= n; i++ {
		// 当前总次数等于上一次和上上次走法之和
		totalSteps = prePre + pre
		// 上上次数变量，赋值为上一次
		prePre = pre
		// 上一次数变量，赋值为本次
		pre = totalSteps
	}
	return totalSteps
}
