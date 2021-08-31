package link

import (
	"bytes"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct{
	Href string
	Text string
}

func LinkFunc(HTML []byte) []Link{
	
	r := bytes.NewReader(HTML)
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
				text := ""
				for sub := n.FirstChild; sub != nil; sub = sub.NextSibling {
					if sub.Type == html.TextNode{
						text += strings.TrimSpace(sub.Data)
					}else if s := sub.FirstChild; sub.Type == html.ElementNode && s != nil {
						text +=  strings.TrimSpace(s.Data)
						}
					text += " "
				}
				current := Link{
					Href: strings.TrimSpace(a.Val),
					Text: strings.TrimSpace(text),
				}
				*links = append(*links, current)
		}}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, links)
	}
}