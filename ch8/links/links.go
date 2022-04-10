package links

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url) // 获取 response
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { // 查看返回的code
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body) // 对 body 进行处理.
	if err != nil {                   // prase 出错了
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	log.Println("Got the page")
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr { // 对他的 Attr 进行遍历.
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				// 这个可以还原出来部分的 URL
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post) // 这里为什么没有用c不提醒呢? ...
	}
	if post != nil {
		post(n)
	}
}

// 如何在这里写单元测试呢?
// func main() {
// 	fmt.Println(Extract("http://qq.com"))
// }
