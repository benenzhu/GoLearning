package main

func main() {
	var s []int
	s = append(s, 1) // push in the stack
	// top := s[len(s)-1] // top of the stack
	s = s[:len(s)-1] // pop
}

// package main

// import "fmt"

// func main() {
// 	var s []int
// 	top := s[len(s)-1]
// 	s[len(s) - 1]
// 	fmt.Println("zhuzhu")
// }
