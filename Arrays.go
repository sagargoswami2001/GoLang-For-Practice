package main

import "fmt"

func main() {
	
	var s [5]int
	fmt.Println("emp:", s)

	s[4] = 100
	fmt.Println("set:", s)
	fmt.Println("get:", s[4])

	fmt.Println("len:", len(s))

	m := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", m)

	var twoD [2] [3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
