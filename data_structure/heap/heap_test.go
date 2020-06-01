package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	list := []int{3, 6, 2, 7, 1, 9, 4}
	sh := NewHeap(len(list), SmallHeap)
	num, ok := sh.Top()
	fmt.Printf("sh top num:%d, ok:%v\n", num, ok)
	for _, num := range list {
		sh.Insert(num)
	}
	num, ok = sh.Top()
	fmt.Printf("after insert, sh top num:%d, ok:%v\n", num, ok)
	fmt.Printf("sh list:%v\n", sh.List())

	topNum, ok := sh.Delete()
	fmt.Printf("sh delete num:%v, ok:%v\n", topNum, ok)
	fmt.Printf("sh list:%v\n", sh.List())
	topNum, ok = sh.Delete()
	fmt.Printf("sh delete num:%v, ok:%v\n", topNum, ok)
	fmt.Printf("sh list:%v\n", sh.List())
	topNum, ok = sh.Delete()
	fmt.Printf("sh delete num:%v, ok:%v\n", topNum, ok)
	fmt.Printf("sh list:%v\n", sh.List())
	topNum, ok = sh.Delete()
	fmt.Printf("sh delete num:%v, ok:%v\n", topNum, ok)
	fmt.Printf("sh list:%v\n", sh.List())

	list1 := []int{7, 9, 8, 16}
	bigHeap := NewHeap(3, BigHeap)
	for _, n := range list1 {
		fmt.Printf("bigHeap insert %d res:%v\n", n, bigHeap.Insert(n))
	}
}
