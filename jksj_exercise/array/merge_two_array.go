package main

func mergeTwoSequentialArray(arr1, arr2 []int) []int {
	size1 := len(arr1)
	size2 := len(arr2)
	if size1 == 0 {
		return arr2
	} else if size2 == 0 {
		return arr1
	}
	newArray := make([]int, size1+size2)
	var i, j, n int
	// compare two arrays' element, put them into newArray
	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			newArray[n] = arr1[i]
			i++
		} else {
			newArray[n] = arr2[j]
			j++
		}
		n++
	}
	// put the rest of elements into newArray, the rest of elements may in arr1 or arr2
	for i < size1 {
		newArray[n] = arr1[i]
		i++
		n++
	}
	for j < size2 {
		newArray[n] = arr2[j]
		j++
		n++
	}
	return newArray
}


