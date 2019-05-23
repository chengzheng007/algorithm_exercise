package main

import "fmt"

func main() {
	hashmap := NewHashMap(10)
	hashmap.Add(1)
	hashmap.Add(2)
	hashmap.Add(3)
	hashmap.Add(5)
	hashmap.Add(8)
	hashmap.Print()

	hashmap.Add(15)
	hashmap.Add(35)

	hashmap.Add(9)
	hashmap.Add(19)
	hashmap.Add(219)
	ret := hashmap.Add(3)
	if !ret {
		fmt.Println("add 3 failed, num existed")
	}
	hashmap.Add(35)
	hashmap.Print()

	hashmap.Del(2)
	fmt.Println("after del 2:")
	hashmap.Print()

	hashmap.Del(5)
	fmt.Println("after del 5:")
	hashmap.Print()

	hashmap.Del(35)
	fmt.Println("after del 35:")
	hashmap.Print()

	hashmap.Del(15)
	fmt.Println("after del 15:")
	hashmap.Print()
}
