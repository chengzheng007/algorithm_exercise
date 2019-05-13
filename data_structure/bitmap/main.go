package main

import "fmt"

func main() {
	bitMap := NewBitMap(10000)
	fmt.Printf("bitmap len:%d\n", bitMap.Len())


	var num int64 = 1
	bitMap.Set(num)

	num = 2
	bitMap.Set(num)

	num = 3
	bitMap.Set(num)

	num = 88
	bitMap.Set(num)

	num = 89
	bitMap.Set(num)
	num = 90
	bitMap.Set(num)
	num = 91
	bitMap.Set(num)
	num = 92
	bitMap.Set(num)
	num = 93
	bitMap.Set(num)
	num = 94
	bitMap.Set(num)
	num = 95
	bitMap.Set(num)


	num = 9900
	bitMap.Set(num)

	num = 192
	fmt.Printf("%d exist? %v\n", num, bitMap.Get(num))

	num = 9900
	fmt.Printf("%d exist? %v\n", num, bitMap.Get(num))

	num = 8888
	fmt.Printf("%d exist? %v\n", num, bitMap.Get(num))

	list := bitMap.List()
	fmt.Printf("nums in bitmap:%v\n", list)
}

type bitMap struct {
	byteArr []byte
	nbit int64
}

func NewBitMap(n int64) *bitMap {
	obj := &bitMap{}
	obj.nbit = n
	obj.byteArr = make([]byte, n/8+1) // every byte store 8 bits, every byte range is [0,255]
	return obj
}

func (bm *bitMap) Set(num int64) {
	if num > bm.nbit {
		return
	}
	idx := int(num / 8)
	bitIdx := uint(num % 8)
	bm.byteArr[idx] |= 1 << bitIdx
	fmt.Printf("bm.byteArr[%d]:%d\n", idx, bm.byteArr[idx])
}

func (bm *bitMap) Get(num int64) bool {
	if num > bm.nbit {
		return false
	}
	idx := int(num / 8)
	bitIdx := uint(num % 8)
	return (bm.byteArr[idx] & (1 << bitIdx)) != 0
}

func (bm *bitMap) Len() int {
	return len(bm.byteArr)
}

func  (bm *bitMap) List() []int {
	size := bm.Len()
	var list []int
	for byteIdx := 0; byteIdx < size; byteIdx++ {
		for bitIdx := 0; bitIdx < 8; bitIdx++ {
			if bm.byteArr[byteIdx] & (1 << uint(bitIdx)) != 0 {
				list = append(list, 8*byteIdx+bitIdx)
			}
		}
	}
	return list
}