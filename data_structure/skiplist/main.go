package main

import (
	"math/rand"
	"time"
	"fmt"
)

// 跳表练习

func main() {
	list := []int{6,3,7,9,1,12,35,233}
	psl := newSkipList()
	for _, num := range list {
		psl.Insert(num)
	}
	psl.Print()

	p := psl.Find(6)
	fmt.Printf("p:%+v\n", p)

	p = psl.Find(12)
	fmt.Printf("p:%+v\n", p)

	p = psl.Find(233)
	fmt.Printf("p:%+v\n", p)

	p = psl.Find(12345)
	fmt.Printf("p:%+v\n", p)

	fmt.Println()
	
	delNum := 9
	ret := psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)
	//psl.Print()

	delNum = 12
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)
	//psl.Print()

	delNum = 3
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)

	delNum = 233
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)

	delNum = 1
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)

	delNum = 6
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)

	delNum = 7
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)

	delNum = 35
	ret = psl.Delete(delNum)
	fmt.Printf("del %d ret:%v\n", delNum, ret)

	psl.Print()
}


const (
	MAX_LEVEL = 7
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Node struct {
	Val int
	MaxLevel int  			// 当前节点所在最大层数
	Next [MAX_LEVEL]*Node
}

type SkipList struct {
	MaxLevel int      // 最大层数
	MaxLevelNodes int // 最大层节点数
	Head *Node
}

func newSkipList() *SkipList {
	return &SkipList{
		Head:&Node{},
	}
}

func randomLevel() int {
	var level = 1
	for i := 1; i < MAX_LEVEL; i++ {
		if random.Intn(1000) % 2 == 1 {
			level++
		}
	}
	return level
}

func (psl *SkipList) Insert(x int) {
	if psl.Head == nil {
		fmt.Println("uninitialized skip list")
		return
	}
	
	var update [MAX_LEVEL]*Node
	level := randomLevel()

	// 初始化每个节点指向头结点
	for i := 0; i < MAX_LEVEL; i++ {
		update[i] = psl.Head
	}

	// 寻找每一层插入位置的前驱节点，记录在update中
	p := psl.Head
	for i := level-1; i >= 0; i--  {
		for p.Next[i] != nil && p.Next[i].Val < x {
			p = p.Next[i]
		}
		update[i] = p
	}

	newNode := &Node{
		Val:x,
		MaxLevel:level,
	}

	// 插入操作
	for i := 0; i < level; i++ {
		newNode.Next[i] = update[i].Next[i]
		update[i].Next[i] = newNode
	}

	// 修正跳表参数
	if psl.MaxLevel < level {
		psl.MaxLevel = level
		psl.MaxLevelNodes = 1
	} else if psl.MaxLevel == level {
		psl.MaxLevelNodes++
	}
}

func (psl *SkipList) Print() {
	if psl == nil || psl.Head == nil {
		fmt.Println("skip list has no data")
		return
	}

	fmt.Printf("skip list max level:%d, max level nodes:%d\n", psl.MaxLevel, psl.MaxLevelNodes)

	for i := psl.MaxLevel-1; i >= 0; i-- {
		fmt.Printf("%d(th) level:\n", i+1)
		p := psl.Head
		j := 1
		for p.Next[i] != nil {
			fmt.Printf("node num:%d, val:%v, max level:%d\n", j, p.Next[i].Val, p.Next[i].MaxLevel)
			p = p.Next[i]
			j++
		}
	}
}

// 查找第一个出现的元素x
func (psl *SkipList) Find(x int) *Node {
	if psl.Head == nil {
		return nil
	}
	p := psl.Head
	for i := psl.MaxLevel-1; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < x {
			p = p.Next[i]
		}
	}

	if p.Next[0] != nil && p.Next[0].Val == x {
		return p.Next[0]
	}

	return nil
}

// 删除第一个出现的的元素
// @return true:元素存在并删除  false:未找到元素
func (psl *SkipList) Delete(x int) bool {
	if psl.Head == nil {
		return false
	}

	var prev [MAX_LEVEL]*Node
	p := psl.Head
	for i := psl.MaxLevel-1; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < x {
			p = p.Next[i]
		}
		prev[i] = p
	}

	// 未找到该元素
	if p.Next[0] == nil || p.Next[0].Val != x {
		return false
	}

	if p.Next[0].MaxLevel == psl.MaxLevel {
		psl.MaxLevelNodes--
	}

	for i := psl.MaxLevel-1; i >= 0; i-- {
		if prev[i].Next[i] != nil && prev[i].Next[i].Val == x {
			prev[i].Next[i] = prev[i].Next[i].Next[i]
		}
	}

	// 如果删除导致原本最上一层节点为0，下一层将变为最大层
	if psl.MaxLevelNodes == 0 {
		// MaxLevel-1层为
		for i := psl.MaxLevel-2; i >= 0; i-- {
			p := psl.Head
			for p.Next[i] != nil {
				psl.MaxLevelNodes++
				p = p.Next[i]
			}

			// 已将某一层更新为最上层
			psl.MaxLevel = i + 1
			if psl.MaxLevelNodes > 0 {
				break
			}
		}
	}
	return true
}