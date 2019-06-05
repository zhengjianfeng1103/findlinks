package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("q")
	fmt.Fprintf(w, "Page = %q\n", url)
	if len(url) == 0 {
		return
	}
	page, err := parse("https://" + url)
	if err != nil {
		fmt.Printf("Error getting page %s %s\n", url, err)
		return
	}
	links := pageLink(nil, page)
	for _, link := range links {
		fmt.Fprintf(w, "Link = %q\n", link)
	}
}

func parse(url string) (*html.Node, error) {
	fmt.Println(url)
	//请求
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}
	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse page")
	}
	return b, nil
}

func pageLink(links []string, n *html.Node) []string {
	//拿到节点的a标签
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			//获取a标签中的属性
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//递归
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLink(links, c)
	}
	return links
}
