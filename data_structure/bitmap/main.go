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

// 参考：https://juejin.im/post/5c4fd2af51882525da267385

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
	// 找到位于bitmap数组中的位置，因为我们是按照一个个byte来划分，想象一个大数N
	// 需要N个bit，我们每8个bit划成一份，首先要算出它在哪一份里边
	// 找到位于哪一份之后，在该份中如何表示？要知道我们的数组每一个元素是一个数，范围是[0, 255]
	// 我们其实要找到它在这一份中位于哪个bit，假设是第x个bit，所在的byte数组下标是k(从0开始)
	// 实际大小num=k*8+x，所以它在第k个byte里面的占据的bit是第x个bit，那么它在第k个元素中占用的
	// 的数就是 1 << x，由于一个byte可以表示多个数，这里只能与arr[k]进行&运算，不然这个byte中的
	// 其他数将会丢失
	idx := int(num / 8)
	// 位置具体的数
	bitIdx := uint(num % 8)
	// 将具体的位置为1（假设某数是n，即第n个bit，bitmap中该数大小实际是 1 << n）
	bm.byteArr[idx] |= 1 << bitIdx
	fmt.Printf("set num:%d, bm.byteArr[%d]:%d\n", num, idx, bm.byteArr[idx])
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