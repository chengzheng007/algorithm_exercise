package main

func bf(m, p string) int {
	sizeOfM := len(m)
	sizeOfP := len(p)
	var i, j int
	for i = 0; i < sizeOfM - sizeOfP; i++ {
		for j = 0; j < sizeOfP; j++ {
			if m[i+j] != p[j] {
				break
			}
		}             	
		if j == sizeOfP {
			return i
		}
	}
	return -1
}
