package link

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct{
	Href string
	Text string
}

func LinkFunc(file string) []Link{
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	r := bytes.NewReader(data)
	doc,err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	links := []Link{}

	dfs(doc, &links)
	return links
}

func dfs (n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				current := Link{
					Href: strings.TrimSpace(a.Val),
					Text: strings.TrimSpace(n.FirstChild.Data),
				}
				*links = append(*links, current)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, links)
	}
}