package main

import (
	"fmt"
	"log"

	"zz/ch8/links"
)

var tokens = make(chan struct{}, 5) // 一次最多5个.
func main() {
	worklist := make(chan []string)
	var n int
	n++
	// go func() {worklist <- os.Args[1:]}()
	go func() { worklist <- []string{"https://qq.com"} }() // 用 qq.com 方便test
	used := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !used[link] {
				used[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
	fmt.Println("zhuzhu")
}

func crawl(url string) []string {
	// fmt.Println(url)
	tokens <- struct{}{} // 共享变量？
	fmt.Println("\tstart  Extract ", url)
	list, err := links.Extract(url)
	fmt.Println("\tfinish Extract ", url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
