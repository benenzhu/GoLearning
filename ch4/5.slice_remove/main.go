package main

import "zz/ch4/use"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	ages := map[string]int{}
	s := []int{6, 7, 8, 9}
	s = remove(s, 2)
	use.U(ages, s)
}
