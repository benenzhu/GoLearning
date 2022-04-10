package main

import (
	"fmt"
)

func main() {
	var x []int
	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Println(i, cap(x), len(x), x)
		fmt.Println("\t", x[:cap(x)])
	}
	var y []int
	// var z []int
	for _, i := range y {
		fmt.Println(i)
	}
}

/*
output:
8 16 9 [0 1 2 3 4 5 6 7 8]
         [0 1 2 3 4 5 6 7 8 0 0 0 0 0 0 0]
9 16 10 [0 1 2 3 4 5 6 7 8 9]
         [0 1 2 3 4 5 6 7 8 9 0 0 0 0 0 0]

*/
